Feature: Browse restaurant menus

  Scenario:
    Given a restaurant "miim"
    When I visit retaurant detail page
    Then there should be "Main Menu" menu
    And there should be "Burgers, Pizza, Tequila, Vodka" category
