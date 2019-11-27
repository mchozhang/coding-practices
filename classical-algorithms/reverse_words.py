# reverse words
# input: ['i', ' ', 'l', 'o', 'v', 'e', ' ', 'y', 'o', 'u']
# output: ['y', 'o', 'u', ' ', 'l', 'o', 'v', 'e', ' ', 'i']


def reverse_words(words):
    """
    reverse words:
    traverse the list, insert character in front of the index of the first space of the new list
    time complexity: O(n)
    space complexity: O(n)
    :param words: list of character, words separated by space
    :return: reversed word list
    """
    new_words = []

    # index of the first space in the new word list
    space_index = -1

    for character in words:
        if character == ' ':
            space_index = 0
            new_words.insert(0, character)
            continue

        # no space yet
        if space_index == -1:
            new_words.append(character)
            continue

        new_words.insert(space_index, character)
        space_index += 1

    return new_words


if __name__ == "__main__":
    # words = ['l', 'o', 'v', 'e']
    words = ['b', 'a', 'b', 'y', ' ', 'i', ' ', 'l', 'o', 'v', 'e', ' ', 'y', 'o', 'u']
    print(reverse_words(words))
