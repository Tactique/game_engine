import os


def find_all(directory, pattern):
    files = []

    def prepend_directory(file_):
        return os.path.join(directory, file_)
    for file_ in map(prepend_directory, os.listdir(directory)):
        if os.path.isdir(file_):
            files.extend(find_all(file_, pattern))
        elif os.path.isfile(file_) and file_.endswith(pattern):
            files.append(file_)
    return files
