import json

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
        retstr = '{"%s": {' % self.__class__.__name__
        for attr in sorted(vars(self).keys()):
            _attr_val = getattr(self, attr)
            if isinstance(_attr_val, str):
                attr_val = "\"%s\"" % (_attr_val,)
            elif isinstance(_attr_val, dict):
                attr_val = json.dumps(_attr_val)
            else:
                attr_val = _attr_val
            retstr += '"%s": %s, ' % (attr, attr_val)
        return retstr[:-2] + "}}"


class BaseDictionary(BaseClass):
    def __init__(self, dictionary={}):
        self.dictionary = dictionary

    @contract.returns(None)
    def __setitem__(self, index, multiplier):
        self.dictionary[index] = multiplier

    @contract.self_accepts(int)
    def __getitem__(self, index):
        return self.dictionary[index]

    @contract.returns(int)
    def __len__(self):
        return len(self.dictionary)

    @contract.returns([tuple])
    def items(self):
        return self.dictionary.items()


class BaseEnum(BaseDictionary):
    @contract.self_accepts(str, int)
    @contract.returns(None)
    def __setitem__(self, index, name):
        self.dictionary[index] = name

    @contract.self_accepts(str)
    @contract.returns(int)
    def __getitem__(self, index):
        return self.dictionary[index]


class BaseMultiplier(BaseDictionary):
    @contract.self_accepts(int, float)
    @contract.returns(None)
    def __setitem__(self, index, multiplier):
        self.dictionary[index] = multiplier

    @contract.self_accepts(int)
    @contract.returns(float)
    def __getitem__(self, index):
        return self.dictionary[index]
