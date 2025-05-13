#coding = 'utf-8'

import sys
import typing
from PyQt5 import QtCore, QtGui, QtWidgets
from PyQt5.QtWidgets import QApplication, QDialog, QWidget, QTableWidgetItem, QPushButton, QVBoxLayout, QMessageBox, QGridLayout, QLabel
import Ui_login
import Ui_main
import Ui_config
import Ui_password
import subprocess
from request_operate import RequestOperate
import requests
from config_util import ConfigUtil
from qt_material import apply_stylesheet

config = ConfigUtil()

def show_tips(constant, tip=0):
    msg_box = QMessageBox()
    if tip == 1:
        msg_box.setIcon(QMessageBox.Warning)
        msg_box.setWindowTitle("警告")
    else:
        msg_box.setIcon(QMessageBox.Information)
        msg_box.setWindowTitle("信息")
    msg_box.setText(constant)

    msg_box.exec_()


class ConfigDialog(QtWidgets.QDialog, Ui_config.Ui_Dialog_setConfig):

    def __init__(self, parent=None):
        super(ConfigDialog, self).__init__(parent)
        self.setupUi(self)
        self.pushButton_setconfig.clicked.connect(self.set_config)
        self.lineEdit_serverip.setText(config.read_ini("server", "ip"))

    def set_config(self):
        server_ip = self.lineEdit_serverip.text()
        config.write_ini("server", "ip", server_ip)
        self.close()


class PasswordDialog(QtWidgets.QDialog, Ui_password.Ui_DialogPassword):
    def __init__(self, parent=None):
        super(PasswordDialog, self).__init__(parent)
        self.setupUi(self)
        self.pushButton_modify.clicked.connect(self.set_password)
        self.pushButton_cancel.clicked.connect(self.cancel)

    def cancel(self):
        self.close()

    def set_password(self):
        password = self.lineEdit_password.text()
        repeat = self.lineEdit_repeat.text()
        if password != repeat:
            show_tips("两次输入的密码不一致，请重新输入！", 1)
            return
        url = "/api/cloud/v1/user"
        data = {
            "password": password
        }
        try:
            resp = RequestOperate(server_ip=config.read_ini("server", "ip")).put(url, headers=self.cookie, data=data)
            if resp.get('code') == 0:
                config.write_ini("user", "password", password)
                self.close()
            else:
                show_tips("密码修改失败", 1)
        except Exception as e:
            print(str(e))


    def do_login_cookie(self, data):
        self.cookie = eval(data)


class LoginDialog(QtWidgets.QDialog, Ui_login.Ui_Dialog):

    submitSingal = QtCore.pyqtSignal(str)

    def __init__(self, parent=None):
        super(LoginDialog, self).__init__(parent)
        self.setupUi(self)
        self.mainWindow = MainWindow()
        self.submitSingal.connect(self.mainWindow.do_login_cookie)
        self.pushButton_login.clicked.connect(self.do_login)
        self.pushButton_config.clicked.connect(self.do_config)
        self.lineEdit_username.setText(config.read_ini("user", "name"))
        self.lineEdit_password.setText(config.read_ini("user", "password"))

    def do_login(self):
        server_ip = config.read_ini("server", "ip")
        if server_ip == "":
            show_tips("请先配置服务器地址", 1)
            return
        url = "/api/cloud/v1/login"
        data = {
            "username": self.lineEdit_username.text(),
            "password": self.lineEdit_password.text()
        }
        try:
            resp = RequestOperate(server_ip=server_ip).post(url, data=data)
            if resp.get('code') == 0:
                self.hide()
                self.mainWindow.show()
                self.submitSingal.emit(str(resp.get('data')))
            else:
                if resp.get('msg') == "User.Disable":
                    show_tips("用户已禁用，请联系管理员启用该用户！", 1)
                else:
                    show_tips("用户名或密码错误", 1)
        except Exception as e:
            print(str(e))

    def do_config(self):
        self.config_dlg = ConfigDialog()
        self.config_dlg.show()


class MainWindow(QtWidgets.QMainWindow, Ui_main.Ui_MainWindow):

    submitSingal1 = QtCore.pyqtSignal(str)

    def __init__(self, parent=None):
        super().__init__(parent)
        self.setupUi(self)
        self.cookie = {}
        self.devices = []
        self.pos_dict = {}


    def initTableWidget(self):
        central_widget = QWidget(self)
        central_widget.setFixedSize(800, 800)
        self.setCentralWidget(central_widget)

        grid_layout = QGridLayout(central_widget)


        for i in range(len(self.devices)):
            widget = QWidget()
            widget.setFixedSize(250, 200)
            widget.setObjectName("myWidget")
            widget.setStyleSheet("QWidget
            vm_name = self.devices[i]["name"]
            status = self.devices[i]["status"]
            ip = self.devices[i]["virtualIp"]
            label1 = QLabel(f'云桌面: {vm_name}')
            label2 = QLabel(f'状态: {status}')
            label3 = QLabel(f'IP: {ip}')
            button1 = QPushButton(f'开机')
            button2 = QPushButton(f'关机')
            button3 = QPushButton(f'重启')
            button4 = QPushButton(f'连接')
            layout = QVBoxLayout()
            layout.addWidget(label1)
            layout.addWidget(label2)
            layout.addWidget(label3)
            layout.addWidget(button1)
            layout.addWidget(button2)
            layout.addWidget(button3)
            layout.addWidget(button4)
            widget.setLayout(layout)
            row = i // 3
            col = i % 3
            self.pos_dict[(row, col)] = i
            grid_layout.addWidget(widget, row, col)
            button1.clicked.connect(self.open_vm1(row, col))
            button2.clicked.connect(self.close_vm1(row, col))
            button3.clicked.connect(self.reset_vm1(row, col))
            button4.clicked.connect(self.connect_vm1(row, col))

    def operate_vm1(self, operate, index):
        data = {
            "id": self.devices[index]["id"],
            "operate": operate
        }
        try:
            url = "/api/cloud/v1/devices/operate"
            resp = RequestOperate(server_ip=config.read_ini("server", "ip")).post(url, data=data, headers=self.cookie)
            if resp.get('code') == 0:
                show_tips("操作云桌面成功")
            else:
                if resp.get('msg') == "User.Disable":
                    show_tips("用户已禁用，请联系管理员启用该用户！", 1)
                else:
                    show_tips("操作失败", 1)
        except Exception as e:
            print(str(e))

    def open_vm1(self, row, col):
        def slot():
            index = self.pos_dict.get((row, col), 0)
            operate = "open"
            self.operate_vm1(operate, index)
            self.load_vms()
        return slot

    def close_vm1(self, row, col):
        def slot():
            index = self.pos_dict.get((row, col), 0)
            operate = "close"
            self.operate_vm1(operate, index)
            self.load_vms()
        return slot

    def reset_vm1(self, row, col):
        def slot():
            index = self.pos_dict.get((row, col), 0)
            operate = "reset"
            self.operate_vm1(operate, index)
            self.load_vms()
        return slot

    def do_modify_pwd(self):
        self.pwd_dialog = PasswordDialog()
        self.pwd_dialog.do_login_cookie(str(self.cookie))
        self.pwd_dialog.show()

    def do_login_cookie(self, data):
        data_dict = eval(data)
        self.cookie = {
            'Cookie': 'accessToken={}; userId={}'.format(data_dict.get("value"), data_dict.get("user_id"))
        }
        self.get_user_profile()
        self.load_vms()
        self.pushButton_modify.clicked.connect(self.do_modify_pwd)

    def load_vms(self):
        try:
            url = "/api/cloud/v1/devices"
            resp = RequestOperate(server_ip=config.read_ini("server", "ip")).get(url, headers=self.cookie)
            if resp.get('code') == 0:
                self.devices.clear()
                self.pos_dict.clear()
                self.devices = resp.get('data')
                self.initTableWidget()
            else:
                if resp.get('msg') == "User.Disable":
                    show_tips("用户已禁用，请联系管理员启用该用户！", 1)
                else:
                    show_tips("加载云桌面失败", 1)
        except Exception as e:
            print(str(e))

    def get_user_profile(self):
        try:
            url = "/api/cloud/v1/user/profile"
            resp = RequestOperate(server_ip=config.read_ini("server", "ip")).get(url, headers=self.cookie)
            if resp.get('code') == 0:
                self.label_login_user.setText("用户：" + resp.get('data').get('username'))
            else:
                show_tips("获取用户信息失败", 1)
        except Exception as e:
            print(str(e))

    def connect_vm1(self, row, col):
        def slot():
            index = self.pos_dict.get((row, col), 0)
            server_ip = config.read_ini("server", "ip")
            vm_ip = self.devices[index]["virtualIp"]
            subprocess.call("mstsc /v:{} /w:1920 /h:1080".format(vm_ip))
        return slot


if __name__ == "__main__":
    apply_stylesheet(QApplication(sys.argv), theme='dark_blue.xml')
    app = QApplication(sys.argv)
    dlg = LoginDialog()
    dlg.show()
    sys.exit(app.exec_())
