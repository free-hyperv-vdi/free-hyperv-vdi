#coding = 'utf-8'

import configparser


class ConfigUtil(object):

    def __init__(self, file_path="conf.ini"):
        self.file_path = file_path
        self.config = configparser.ConfigParser()

    def read_ini(self, section, key):
        self.config.read(self.file_path)
        try:
            return self.config.get(section, key)
        except:
            return ""

    def write_ini(self, section, key, value):
        if not self.config.has_section(section):
            self.config.add_section(section)
        self.config.set(section, key, value)
        with open(self.file_path, 'w') as config_file:
            self.config.write(config_file)
