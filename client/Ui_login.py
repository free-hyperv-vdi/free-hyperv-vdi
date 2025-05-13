# -*- coding: utf-8 -*-



from PyQt5 import QtCore, QtGui, QtWidgets


class Ui_Dialog(object):
    def setupUi(self, Dialog):
        Dialog.setObjectName("Dialog")
        Dialog.resize(400, 300)
        self.label_username = QtWidgets.QLabel(Dialog)
        self.label_username.setGeometry(QtCore.QRect(80, 80, 54, 12))
        self.label_username.setObjectName("label_username")
        self.label_password = QtWidgets.QLabel(Dialog)
        self.label_password.setGeometry(QtCore.QRect(80, 130, 54, 12))
        self.label_password.setObjectName("label_password")
        self.lineEdit_username = QtWidgets.QLineEdit(Dialog)
        self.lineEdit_username.setGeometry(QtCore.QRect(150, 80, 161, 20))
        self.lineEdit_username.setObjectName("lineEdit_username")
        self.pushButton_login = QtWidgets.QPushButton(Dialog)
        self.pushButton_login.setGeometry(QtCore.QRect(90, 200, 111, 23))
        self.pushButton_login.setObjectName("pushButton_login")
        self.lineEdit_password = QtWidgets.QLineEdit(Dialog)
        self.lineEdit_password.setGeometry(QtCore.QRect(150, 130, 161, 20))
        self.lineEdit_password.setEchoMode(QtWidgets.QLineEdit.Password)
        self.lineEdit_password.setObjectName("lineEdit_password")
        self.pushButton_config = QtWidgets.QPushButton(Dialog)
        self.pushButton_config.setGeometry(QtCore.QRect(234, 200, 111, 23))
        self.pushButton_config.setObjectName("pushButton_config")

        self.retranslateUi(Dialog)
        QtCore.QMetaObject.connectSlotsByName(Dialog)

    def retranslateUi(self, Dialog):
        _translate = QtCore.QCoreApplication.translate
        Dialog.setWindowTitle(_translate("Dialog", "云平台客户端"))
        self.label_username.setText(_translate("Dialog", "用户名："))
        self.label_password.setText(_translate("Dialog", "密 码："))
        self.pushButton_login.setText(_translate("Dialog", "登录"))
        self.pushButton_config.setText(_translate("Dialog", "配置"))
