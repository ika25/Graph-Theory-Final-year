# Graph-Theory-Final-year
This is a program written in Go language that can build a non-deterministic ﬁnite automaton (NDFA) from a regular expression, and can use the NFA to check if the regular expression matches any given string of text.

# How to run program:
1)Install Go Lang in your PC

2)Install GO Programming in VS code: https://code.visualstudio.com/docs/languages/go

3)Start Deuging in VS code and that should run your project.

# what is ndfa?
In NDFA, for a particular input symbol, the machine can move to any combination of the states in the machine. In other words, the exact state to which the machine moves cannot be determined. Hence, it is called Non-deterministic Automaton. As it has finite number of states, the machine is called Non-deterministic Finite Machine or Non-deterministic Finite Automaton.

# Formal Definition of an NDFA
An NDFA can be represented by a 5-tuple (Q, ∑, δ, q0, F) where −

Q is a finite set of states.

∑ is a finite set of symbols called the alphabets.

δ is the transition function where δ: Q × ∑ → 2Q

(Here the power set of Q (2Q) has been taken because in case of NDFA, from a state, transition can occur to any combination of Q states)

q0 is the initial state from where any input is processed (q0 ∈ Q).

F is a set of final state/states of Q (F ⊆ Q).

# Graphical Representation of an NDFA:
An NDFA is represented by digraphs called state diagram.

The vertices represent the states.

The arcs labeled with an input alphabet show the transitions.

The initial state is denoted by an empty single incoming arc.

The final state is indicated by double circles.

# This program uses two algorithm
1)Thompson's Construction Algorithm - https://en.wikipedia.org/wiki/Thompson%27s_construction

2)Shunt Algorithm - https://en.wikipedia.org/wiki/Shunting-yard_algorithm

# References
https://code.visualstudio.com/docs/editor/versioncontrol

StackOverFlow

Instructions provided by Dr Ian McLoughlin - https://web.microsoftstream.com/video/68a288f5-4688-4b3a-980e-1fcd5dd2a53b, https://web.microsoftstream.com/video/bad665ee-3417-4350-9d31-6db35cf5f80d
