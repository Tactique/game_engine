use_color = True

red = '31'
green = '32'
blue = '34'


def colorize(string, color):
    if use_color:
        return '\33[%sm%s\33[0m' % (color, string)
    else:
        return string
