def removed(list_, to_remove):
    def no_element(element):
        return element != to_remove

    return filter(no_element, list_)
