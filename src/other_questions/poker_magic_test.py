# poker magic test


def generate_cards(suit_list, start, interval, total_card=52):
    """
    generate pack of cards by the specified rule
    :param suit_list: suit list
    :param start: start number
    :param interval: card number interval
    :param total_card: total number of card
    :return: list of card
    """
    number_list = ['A', '2', '3', '4', '5', '6', '7', '8', '9', '10', 'J', 'Q', 'K']
    cards = []
    i = 0
    current = start
    while i < total_card:
        number = number_list[current]
        suit = suit_list[i % 4]
        cards.append((number, suit))

        current += interval
        if current >= len(number_list):
            current -= len(number_list)
        i += 1

    return cards


def shuffle(cards, index):
    """
    put the cards of range 0 - index to the bottom
    :param cards: list of card
    :param index: selected index (0 - last)
    :return: new card list
    """
    return cards[index:] + cards[:index]


if __name__ == "__main__":
    suits = ['♠', '♥', '♣', '♦']
    pack = generate_cards(suits, 2, 3)

    for index, card in enumerate(pack):
        print(index, card)
    print("shuffle")

    pack = shuffle(pack, 10)
    pack = shuffle(pack, 10)
    pack = shuffle(pack, 12)

    for index, card in enumerate(pack):
        print(index, card)
