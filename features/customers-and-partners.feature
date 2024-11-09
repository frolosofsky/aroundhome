Feature: Customer communication with partners

  # Note, the city names are completely ignored by the test system. They are provided only for a reference.
  # Expected distance is calculated on https://latlongdata.com/distance-calculator/
  # Coords are in the lat, long format.
  Background:
    Given partners:
      | name   | experience    | location           | city      | radius | rating |
      | alice  | wood          | 52.52357, 13.38414 | Berlin    | 10 km  | 9      | 
      | bob    | wood, carpet  | 52.40612, 12.55835 | Magdeburg | 200 km | 8      |
      | carlos | carpet        | 54.08988, 12.13590 | Rostok    | 100 km | 8      |
      | chad   | wood, tiles   | 51.33843, 12.38838 | Leipzig   | 80 km  | 2      |

  Scenario: Customer finds matching partners, results sorted by rating
    When a customer located in "Berlin" ("52.50950", "13.43997") searches someone experienced with "wood"
    Then they should receive a list of partners:
      | name  | rating | distance |
      | alice | 9      | 4 km     |
      | bob   | 8      | 60 km    |

  Scenario: Customer finds matching partners, results sorted by distance after the rating
    When a customer located in "Wismar" ("53.89759", "11.46165") searches someone experienced with "carpet"
    Then they should receive a list of partners:
      | name   | rating | distance |
      | carlos | 8      | 49 km    |
      | bob    | 8      | 181 km   |

  Scenario: Customer on the North Pole cannot find a parter
    When a customer located in "North Pole" ("89.77623", "49.35515") searches someone experienced with "tiles"
    Then they should receive an empty list of partners

  Scenario: Customer get details about a known partner
    When a customer asks details about partner "alice"
    Then they should receive parter details as such
    | name   | experience    | location           | city      | radius | rating |
    | alice  | wood          | 52.52357, 13.38414 | Berlin    | 10     | 9      | 

  Scenario: Customer doesn't receive details of unknown partner
    When a customer asks details about partner "craig"
    Then they should receive error (404)
