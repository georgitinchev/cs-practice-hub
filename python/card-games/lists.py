"""Functions for tracking poker hands and assorted card tasks.

Python list documentation: https://docs.python.org/3/tutorial/datastructures.html
"""
import statistics

def get_rounds(number):
    """Create a list containing the current and next two round numbers.

    :param number: int - current round number.
    :return: list - current round and the two that follow.
    """
    arr = list((number,))
    arr.append(number + 1)
    arr.append(number + 2)
    return arr



def concatenate_rounds(rounds_1, rounds_2):
    """Concatenate two lists of round numbers.

    :param rounds_1: list - first rounds played.
    :param rounds_2: list - second set of rounds played.
    :return: list - all rounds played.
    """
    arr = rounds_1 + rounds_2
    return arr


def list_contains_round(rounds, number):
    """Check if the list of rounds contains the specified number.

    :param rounds: list - rounds played.
    :param number: int - round number.
    :return: bool - was the round played?
    """

    for round in rounds:
        if round == number:
            return True
    return False

def card_average(hand):
    """Calculate and returns the average card value from the list.

    :param hand: list - cards in hand.
    :return: float - average value of the cards in the hand.
    """
    sumHand = sum(hand)
    average = sumHand / len(hand)
    return float(average)


def approx_average_is_average(hand):
    """Return if the (average of first and last card values) OR ('middle' card) == calculated average.

    :param hand: list - cards in hand.
    :return: bool - does one of the approximate averages equal the `true average`?
    """
    actualAvg = card_average(hand)
    firstAvg = (hand[0] + hand[-1]) / 2
    secondAvg = float(statistics.median(hand))
    if actualAvg == firstAvg or actualAvg == secondAvg:
        return True
    else:
        return False
    


def average_even_is_average_odd(hand):
    """Return if the (average of even indexed card values) == (average of odd indexed card values).
    :param hand: list - cards in hand.
    :return: bool - are even and odd averages equal?
    """
    evenAvgs = card_average(hand[::2]) 
    oddAvgs = card_average(hand[1::2])
    if evenAvgs == oddAvgs:
        return True
    else:
        return False

def maybe_double_last(hand):
    """Multiply a Jack card value in the last index position by 2.

    :param hand: list - cards in hand.
    :return: list - hand with Jacks (if present) value doubled.
    """
    arr = hand[:]  
    if arr[-1] == 11:
        arr[-1] = arr[-1] * 2
    return arr
    
