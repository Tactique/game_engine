import types


class ContractBrokenError(Exception):
    def __init__(self, message):
        Exception.__init__(self, message)


def returns(return_type):
    def _returns(orig_func):
        def helper(*args, **kwargs):
            ret = orig_func(*args, **kwargs)
            if return_type is None:
                if ret is not None:
                    raise ContractBrokenError(
                        'Type %s does not match contract type %s' % (
                            type(ret), return_type))
            else:
                _assert_is_type(ret, return_type)
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
    if not isinstance(arg, contract_type):
        raise ContractBrokenError(
            'Type %s does not match contract type %s' % (
                arg_type, contract_type))


def _assert_same_length(args, contract_args):
    message = 'Number of arguments (%s) does not match contract (%s))'
    if len(args) != len(contract_args):
        raise ContractBrokenError(message % (args, contract_args))
