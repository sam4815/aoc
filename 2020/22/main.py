from time import time

start = time()

with open("input.txt", encoding="utf8") as file:
    regular_decks = [[int(x) for x in deck.splitlines()[1:]] for deck in file.read().strip().split("\n\n")]
    recursive_decks = [deck[:] for deck in regular_decks]


def play_recursive_round(decks):
    card_a, card_b = decks[0].pop(0), decks[1].pop(0)

    if len(decks[0]) >= card_a and len(decks[1]) >= card_b:
        sub_decks = [decks[0][:card_a], decks[1][:card_b]]
        round_num = 0

        while len(sub_decks[0]) > 0 and len(sub_decks[1]) > 0:
            play_recursive_round(sub_decks)

            round_num += 1
            if round_num > 200:
                decks[0] += [card_a, card_b]
                return

        if len(sub_decks[0]) > 0:
            decks[0] += [card_a, card_b]
        else:
            decks[1] += [card_b, card_a]

    elif card_a > card_b:
        decks[0] += [card_a, card_b]

    else:
        decks[1] += [card_b, card_a]


def play_regular_round(decks):
    card_a, card_b = decks[0].pop(0), decks[1].pop(0)

    if card_a > card_b:
        decks[0] += [card_a, card_b]
    else:
        decks[1] += [card_b, card_a]


def find_winning_score(decks, play):
    while len(decks[0]) > 0 and len(decks[1]) > 0:
        play(decks)

    winning_deck = decks[0] if len(decks[0]) > 0 else decks[1]
    return sum((i + 1) * x for i, x in enumerate(winning_deck[::-1]))


winning_regular_score = find_winning_score(regular_decks, play_regular_round)
winning_recursive_score = find_winning_score(recursive_decks, play_recursive_round)

time_elapsed = round(time() - start, 5)

print(
    f"""The winning player's score in regular Combat is {winning_regular_score}.
The winning player's score in recursive Combat is {winning_recursive_score}.
Solution generated in {time_elapsed}s."""
)
