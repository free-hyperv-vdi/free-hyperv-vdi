# -*- coding: utf-8 -*-



from PyQt5 import QtCore, QtGui, QtWidgets


class Ui_Dialog_setConfig(object):
    def setupUi(self, Dialog_setConfig):
        Dialog_setConfig.setObjectName("Dialog_setConfig")
        Dialog_setConfig.resize(400, 300)
        self.label = QtWidgets.QLabel(Dialog_setConfig)
        self.label.setGeometry(QtCore.QRect(30, 30, 81, 20))
        self.label.setObjectName("label")
        self.lineEdit_serverip = QtWidgets.QLineEdit(Dialog_setConfig)
        self.lineEdit_serverip.setGeometry(QtCore.QRect(130, 30, 113, 20))
        self.lineEdit_serverip.setObjectName("lineEdit_serverip")
        self.pushButton_setconfig = QtWidgets.QPushButton(Dialog_setConfig)
        self.pushButton_setconfig.setGeometry(QtCore.QRect(130, 210, 75, 23))
        self.pushButton_setconfig.setObjectName("pushButton_setconfig")

        self.retranslateUi(Dialog_setConfig)
        QtCore.QMetaObject.connectSlotsByName(Dialog_setConfig)

    def retranslateUi(self, Dialog_setConfig):
        _translate = QtCore.QCoreApplication.translate
        Dialog_setConfig.setWindowTitle(_translate("Dialog_setConfig", "配置"))
        self.label.setText(_translate("Dialog_setConfig", "服务器地址："))
        self.pushButton_setconfig.setText(_translate("Dialog_setConfig", "配置"))
