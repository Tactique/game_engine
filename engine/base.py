from lib import contract


class BaseClass(object):
    @contract.returns(bool)
    def __ne__(self, base_class):
        return not (self == base_class)
        #return not self.__eq__(base_class)


class BaseDictionary(BaseClass):
    @contract.returns(bool)
    def __eq__(self, base_dictionary):
        return (
            type(self) == type(base_dictionary) and
            self.dictionary == base_dictionary.dictionary)

    @contract.self_accepts(int, float)
    @contract.returns(None)
    def __setitem__(self, index, multiplier):
        self.dictionary[index] = multiplier

    @contract.self_accepts(int)
    @contract.returns(float)
    def __getitem__(self, index):
        return self.dictionary[index]
