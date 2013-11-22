def removed(list_, to_remove):
    def no_element(element):
        return element != to_remove

    return filter(no_element, list_)


def multi_map(func, arg_list):
    return [func(*args) for args in arg_list]
