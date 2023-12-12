# Guide

This guide and codebase is a companion resource for those reading [On the Building Blocks of Mathematical Logic](https://content.wolfram.com/uploads/sites/43/2020/12/Schonfinkel-OnTheBuildingBlocksOfMathematicalLogic.pdf) by [Moses Schönfinkel](https://en.wikipedia.org/wiki/Moses_Sch%C3%B6nfinkel).

The guide annotates the paper section-by-section. I only quote directly from the paper when there is something to call out explicitly. When possible I bias towards a fully working implementation in code that the reader can use themselves.

## Introduction by Quine

The English translation (translated by Stefan Bauer-Mengelberg) of this paper kicks off with an introduction by [Willard Van Orman Quine](https://en.wikipedia.org/wiki/Willard_Van_Orman_Quine). Quine is a philosopher/logician, but you might know him for the type of computer program named after him, the [quine](https://en.wikipedia.org/wiki/Quine_(computing)).

Quine explains the aim of the paper well:

> The initial aim of the paper is reduction of the number of primitive notions of logic. [...] Then presently the purpose deepens. It comes to be nothing less than the general elimination of variables

So Schönfinkel is attempting to simplify logic as far as it can go. Quine explains that there is some prior art here, namely the [elimination of quantifiers](https://en.wikipedia.org/wiki/Quantifier_elimination) (which will be explained in [section 1](./GUIDE.md#section-1), don't worry about the logic formulas for now).

He goes on to explain the "crux" of how to achieve this simplicity:

> The crux of the matter is that Schönfinkel lets functions stand as arguments.

The next paragraph is a brief introduction to [currying](https://en.wikipedia.org/wiki/Currying) (named after [Haskell Curry](https://en.wikipedia.org/wiki/Haskell_Curry), but first discovered by [Gottlob Frege](https://en.wikipedia.org/wiki/Gottlob_Frege)), which we will explain in detail in [section 2](./GUIDE.md#section-2).

The next paragraph tackles the core of the paper, and I think it is quite well-said. Here is the summary:
1. Schönfinkel's calculus assumes on operation: function application
1. He also assumes three specific functions:
   1. `U` - A [universally quantified](https://en.wikipedia.org/wiki/Universal_quantification) [Sheffer stroke](https://en.wikipedia.org/wiki/Sheffer_stroke) (more on this in [section 1](./GUIDE.md#section-1))
   1. `C` - The constancy function (in modern times known as `K`)
   1. `S` - The fusion function
1. Any sentence in logic (including ones with quantification) can be translated into a sentence including just `U`, `C`, `S`, and whatever free variables the original sentence had.

I think this is pretty awesome. Three foundational functions with function application can give you all of logic.

Quine makes a note about the fact that we still technically have variables in the form of a function's definition (ex: `Cxy => x`), but explains:

> But here the variables may be seen as schematic letters for metalogical exposition

This doesn't seem too rigorous to me, but [metalogic](https://en.wikipedia.org/wiki/Metalogic) is a little above my paygrade so I'll take it at face value.

Quine then explains that `U` can actually be derived from `S` and `C` as well as the identity combinator, `I`. On this topic, Quine explains that Schönfinkel attempts to actually reduce our set of foundational functions down to just `J`, from which `S`, `C`, and `U` can be regained. This topic intersects with the modern [Iota](https://en.wikipedia.org/wiki/Iota_and_Jot) combinator.

The next two paragraphs I personally find very interesting. When first reading about Combinatory Logic, I always found the parentheses (essentially the tree structure) to be a heavy assumption that should be given more attention. The background here (that we will learn later) is that Combinatory Logic is [left-associative](https://en.wikipedia.org/wiki/Operator_associativity) (if no parentheses are given they are assumed to be `(((x)x)x)` for `xxx`, for example). Quine's two stated approaches to eliminating the parentheses are:
1. Design the foundational functions all parentheses are "maneuvered" into a left-associative position.
1. Include an operator `o` which is responsible for embedding the parentheses (essentially encoding [Łukasiewicz, or Polish notation](https://en.wikipedia.org/wiki/Polish_notation)).

We will talk about this topic later in [section 6](./GUIDE.md#section-6).

Before switching to the discussion of how to axiomatize this system, Quine briefly remarks about the philosophy of reduction, which is quite interesting (how far can we go to reduce the primitive notations of this system?). He explains that the axiomatization of this system was done by Haskell Curry, under the name of [Combinatory Logic](https://en.wikipedia.org/wiki/Combinatory_logic).

The rest of the introduction is a summary/conclusion to what has been discussed above.

It is at this point that I'd like to point out some additional resources related to Combinatory Logic that I've found particularly helpful:

- [Stanford's Plato article on Combiantory Logic](https://plato.stanford.edu/entries/logic-combinatory/)
- [Stephen Wolfram](https://en.wikipedia.org/wiki/Stephen_Wolfram)'s deep dives into the subject (and the history of Moses Schönfinkel himself)
  - [Combinators and the Story of Computation](https://writings.stephenwolfram.com/2020/12/combinators-and-the-story-of-computation/)
  - [Where Did Combinators Come From? Hunting the Story of Moses Schönfinkel](https://writings.stephenwolfram.com/2020/12/where-did-combinators-come-from-hunting-the-story-of-moses-schonfinkel/)
  - [Combinators: A Centennial View](https://writings.stephenwolfram.com/2020/12/combinators-a-centennial-view/)
- [Raymond Smullyan](https://en.wikipedia.org/wiki/Raymond_Smullyan)'s [To Mock a Mocking Bird](https://en.wikipedia.org/wiki/To_Mock_a_Mockingbird)
- [Ben Lynn](https://crypto.stanford.edu/~blynn/)'s wonderful [website](https://crypto.stanford.edu/~blynn/lambda/sk.html)

## Section 1

## Section 2

## Section 3

## Section 4

## Section 5

## Section 6

## Overall Thoughts
1. Sometimes you can reduce logic down so far that you are actually going in the other direction of simplicity.