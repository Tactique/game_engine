import color


class Notifier:
    def __init__(self, name, allow_failure=False):
        self.name = name
        self.allow_failure = allow_failure

    def success(self, message):
        print _notify(True, self.name, message)

    def failure(self, message):
        print _notify(False, self.name, message)
        if not self.allow_failure:
            raise Exception(_notify(False, self.name, message))


def trail_spaces(string, length):
    return string + (' ' * (length - len(string)))


def _notify(ok, title, message):
    status = color.colorize('OK', color.green) if ok else color.colorize('FAIL', color.red)
    return '%s: %s | %s' % (
        trail_spaces(color.colorize(title.title(), color.blue), 15),
        status,
        message)
