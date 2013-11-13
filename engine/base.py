from lib import contract


class BaseClass(object):
    @contract.returns(bool)
    def __eq__(self, base_class):
        if type(self) != type(base_class):
            return False
        for self_attr, other_attr in zip(vars(self).keys(), vars(base_class).keys()):
            if getattr(self, self_attr) != getattr(base_class, other_attr):
                return False
        return True

    @contract.returns(bool)
    def __ne__(self, base_class):
        return not self.__eq__(base_class)



class BaseDictionary(BaseClass):
    @contract.self_accepts(int, float)
    @contract.returns(None)
    def __setitem__(self, index, multiplier):
        self.dictionary[index] = multiplier

    @contract.self_accepts(int)
    @contract.returns(float)
    def __getitem__(self, index):
        return self.dictionary[index]
