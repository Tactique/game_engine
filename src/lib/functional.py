def removed(list_, to_remove):
    def no_element(element):
        return element != to_remove

    return filter(no_element, list_)


def multi_map(func, arg_list):
    return [func(*args) for args in arg_list]


def multi_reduce(func, arg_list, last_value):
    for args in arg_list:
        last_value = func(last_value, *args)
    return last_value
