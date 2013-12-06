import types


class ContractBrokenError(Exception):
    def __init__(self, message):
        Exception.__init__(self, message)


def returns(return_type):
    def _returns(orig_func):
        def helper(*args, **kwargs):
            ret = orig_func(*args, **kwargs)
            ret_type = return_type
            if return_type is None:
                ret_type = types.NoneType
            _assert_is_type(ret, ret_type)
            return ret
        return helper
    return _returns


def accepts(*arg_types):
    return _accept_base_method(False, *arg_types)


def self_accepts(*arg_types):
    return _accept_base_method(True, *arg_types)


def _accept_base_method(self_or_not, *orig_arg_types):
    def _accepts(orig_func):
        def helper(*args):
            arg_types = list(orig_arg_types)
            if self_or_not:
                arg_types.insert(0, type(args[0]))
            _assert_same_length(args, arg_types)
            for arg, arg_type in zip(args, arg_types):
                _assert_is_type(arg, arg_type)
            return orig_func(*args)
        return helper
    return _accepts


def _assert_is_type(arg, contract_type):
    arg_type = type(arg)
    if isinstance(contract_type, list) or isinstance(contract_type, tuple):
        for sub_arg in arg:
            _assert_is_type(sub_arg, contract_type[0])
        contract_type = type(contract_type)
    if isinstance(contract_type, dict):
        for key, val in arg.items():
            _assert_is_type(key, contract_type.keys()[0])
            _assert_is_type(val, contract_type.values()[0])
        contract_type = type(contract_type)
    if not isinstance(arg, contract_type):
        raise ContractBrokenError(
            'Type %s does not match contract type %s' % (
                arg_type, contract_type))


def _assert_same_length(args, contract_args):
    message = 'Number of arguments (%s) does not match contract (%s))'
    if len(args) != len(contract_args):
        raise ContractBrokenError(message % (args, contract_args))
