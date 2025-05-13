# -*- coding: utf-8 -*-



from PyQt5 import QtCore, QtGui, QtWidgets


class Ui_DialogPassword(object):
    def setupUi(self, DialogPassword):
        DialogPassword.setObjectName("DialogPassword")
        DialogPassword.resize(400, 300)
        self.label = QtWidgets.QLabel(DialogPassword)
        self.label.setGeometry(QtCore.QRect(70, 40, 54, 12))
        self.label.setObjectName("label")
        self.lineEdit_password = QtWidgets.QLineEdit(DialogPassword)
        self.lineEdit_password.setGeometry(QtCore.QRect(160, 40, 113, 20))
        self.lineEdit_password.setObjectName("lineEdit_password")
        self.lineEdit_repeat = QtWidgets.QLineEdit(DialogPassword)
        self.lineEdit_repeat.setGeometry(QtCore.QRect(160, 100, 113, 20))
        self.lineEdit_repeat.setObjectName("lineEdit_repeat")
        self.label_2 = QtWidgets.QLabel(DialogPassword)
        self.label_2.setGeometry(QtCore.QRect(50, 100, 71, 20))
        self.label_2.setObjectName("label_2")
        self.pushButton_modify = QtWidgets.QPushButton(DialogPassword)
        self.pushButton_modify.setGeometry(QtCore.QRect(70, 210, 75, 23))
        self.pushButton_modify.setObjectName("pushButton_modify")
        self.pushButton_cancel = QtWidgets.QPushButton(DialogPassword)
        self.pushButton_cancel.setGeometry(QtCore.QRect(200, 210, 75, 23))
        self.pushButton_cancel.setObjectName("pushButton_cancel")

        self.retranslateUi(DialogPassword)
        QtCore.QMetaObject.connectSlotsByName(DialogPassword)

    def retranslateUi(self, DialogPassword):
        _translate = QtCore.QCoreApplication.translate
        DialogPassword.setWindowTitle(_translate("DialogPassword", "密码修改（云机入口密码）"))
        self.label.setText(_translate("DialogPassword", "新密码："))
        self.label_2.setText(_translate("DialogPassword", "重复新密码："))
        self.pushButton_modify.setText(_translate("DialogPassword", "修改"))
        self.pushButton_cancel.setText(_translate("DialogPassword", "取消"))
