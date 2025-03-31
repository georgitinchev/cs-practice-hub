"""Functions used in preparing Guido's gorgeous lasagna.

Learn about Guido, the creator of the Python language:
https://en.wikipedia.org/wiki/Guido_van_Rossum

This is a module docstring, used to describe the functionality
of a module and its functions and/or classes.
"""


EXPECTED_BAKE_TIME = 40 
PREPARATION_TIME = 2 

def bake_time_remaining(elapsed_bake_time):
    ''' Calculates the remaining baking time
        :param elapsed_bake_time: int - elapsed cooking time

        This function subtracts the elapsed bake time from the expected bake time constant and returns the difference
    '''
    remaining_bake_time = EXPECTED_BAKE_TIME - elapsed_bake_time
    if remaining_bake_time >= 0:
        return remaining_bake_time
    else:
        return 0

def preparation_time_in_minutes(number_of_layers):
    ''' Calculates the preparation time 
        :param number_of_layers: int - number of lasagna layers
        :return: int - Preparation time in integer for set layers derived by multiplying the layers by the preparation time constant
    '''
    if number_of_layers > 0:
        return number_of_layers * PREPARATION_TIME
    else:
        return 0

def elapsed_time_in_minutes(number_of_layers,elapsed_bake_time):
    ''' Returns elapsed overall prep time 
        :param number_of_layers: int - the number of layers in the lasagna
        :param elapsed_bake_time: int - elapsed cooking time
        :return: int - total time elapsed preparing and cooking
    '''
    prep_time = number_of_layers * PREPARATION_TIME
    return prep_time + elapsed_bake_time
