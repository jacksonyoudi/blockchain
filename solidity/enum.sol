pragma solidity ^0.4.0;


contract test {

    enum ActionChoices {GoLeft, GoRight, GoStraight, StiStill}

    ActionChoices choice;

    ActionChoices  constant defaultChoice = ActionChoices.GoStraight;

    function setGoStraight() {
        choice = ActionChoices.GoStraight;
    }

    function getChoice() returns (ActionChoices) {
        return choice;
    }
}



