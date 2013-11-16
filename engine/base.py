from lib import contract


class BaseClass(object):
    @contract.returns(bool)
    def __eq__(self, base_class):
        if type(self) != type(base_class):
            return False
        for self_attr, other_attr in zip(vars(self).values(), vars(base_class).values()):
            if self_attr != other_attr:
                return False
        return True

    @contract.returns(bool)
    def __ne__(self, base_class):
        return not self.__eq__(base_class)

    @contract.returns(str)
    def __repr__(self):
        retstr = ''
        for attr in sorted(vars(self).keys()):
            retstr += '%s:%s ' % (attr, getattr(self, attr))
        return retstr


class BaseDictionary(BaseClass):
    @contract.self_accepts(int, float)
    @contract.returns(None)
    def __setitem__(self, index, multiplier):
        self.dictionary[index] = multiplier

    @contract.self_accepts(int)
    @contract.returns(float)
    def __getitem__(self, index):
        return self.dictionary[index]
