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

### Sheffer Stroke

The first paragraph really sets the tone for Schönfinkel's goal for the rest of the paper. Not only do we want to reduce the number and the content of axioms for a system, but also the *number of fundamental undefined notions*.

Schönfinkel starts out by listing the fundamental propositional connectives,
- $\overline{a}$ (not $a$)
- $a \vee b$ ($a$ or $b$)
- $a \\& b$ ($a$ and $b$)
- $a → b$ (if $a$, then $b$)
- $a \sim b$ ($a$ is equivalent to $b$)

and that there isn't one listed from which all the others can be rederived. There is, though, a particular connective (first published by [Sheffer](https://en.wikipedia.org/wiki/Sheffer_stroke), named the [Sheffer Stroke](https://en.wikipedia.org/wiki/Sheffer_stroke)) that is not in the list above,

$$\overline{a} \vee \overline{b} \\;\\;\\; \text{and} \\;\\;\\; \overline{a \\& b}$$

both of which are the same (using [De Morgan's theorem](https://en.wikipedia.org/wiki/De_Morgan%27s_laws)). Schönfinkel adopts the $|$ symbol as his operator (so $a|b$), and derives the five former fundamental connectives.

### Universally quantified Sheffer stroke

Schönfinkel now wants to move on from [propositional logic](https://en.wikipedia.org/wiki/Propositional_calculus) to [first-order logic](https://en.wikipedia.org/wiki/First-order_logic) to see if he can modify his fundamental connective to include [quantification](https://en.wikipedia.org/wiki/Quantifier_(logic)). This connective he defines as

$$(x)[\overline{f(x)} \vee \overline{g(x)}]$$

or "for all $x$, either $f(x)$ is false or $g(x)$ is false". He extends the Sheffer Stroke symbol $|$ with a [universal quantifier](https://en.wikipedia.org/wiki/Universal_quantification) $x$, resulting in $|^x$. The full connective being

$$f(x) \\;|^x \\;g(x)$$

Schönfinkel then uses this definition to rederive all connectives and quantifiers in first-order logic. It might not be immediately clear what he is doing here. Let's take the $\overline{a}$ as an example. Schönfinkel is saying the $x$ in $a|^xa$ does not play a role, so really we are saying "for all $x$, either $a$ is false or $a$ is false". From this we can be sure that $a$ is false, or $\overline{a}$. He uses similar tactics in the rest of the derivations.

The final two paragraphs of this section expound a bit more on the motivation of this paper. We are motivated not only from a methodological point of view, but also an aesthetic one. In the final sentence Schönfinkel not only tells us he will succeed in his goals but that he finds that he is able to do so "remarkable".

## Section 2

### Higher Order Functions

In this section Schönfinkel gives his definition of a function. He beings by,

> permitting functions themselves to appear as argument values and also as function values

In other words, [higher-order functions](https://en.wikipedia.org/wiki/Higher-order_function) (which as far as I can tell, [Frege](https://softwareengineering.stackexchange.com/questions/186035/who-first-coined-the-term-higher-order-function-and-or-first-class-citizen) was the first to introduce). The application of an argument $x$ to a function $f$ is just:

$$fx$$

### Currying

Next we learn that Schönfinkel wants all of his functions to have only one argument. He is of course refering to [currying](https://en.wikipedia.org/wiki/Currying) (also anticipated by Frege). I think Schönfinkel describes them quite well so I will move on to his final thoughts in this section.

### Left Associativity and Implementation

We know finally have a window into what our functions actually look like:

$$fxyz...$$
$$\text{or}$$
$$(((fx)y)z)...$$

We can leave the parentheses out if we want, by default they are left-associative. We can also add parentheses to make it clear when expressions should not be left-associative, for example

$$f(xy)z$$

It is here where we can begin our implementation. Because it is clear that our expressions have [terms](https://en.wikipedia.org/wiki/Term_(logic)), we will have to represent them as a [tree structure](https://en.wikipedia.org/wiki/Tree_(data_structure)).

The tree structure representing our expression of functions is peculiar in that:
1. Each function has only one argument
1. Functions associate to the left
1. Our expressions will not have "variables"

Given the expression above, $f(xy)z$, we will represent the tree as the following:

```
f(xy)z
------

  /\
 /\ z
f /\
 x  y
```

Here are a few more trees to get us started:
```
x
-

x
```

```
xy
--

 /\
x  y
```

```
xyz or (xy)z
------------

  /\
 /\ z
x  y
```

This structure
1. Allows us to swap in-place expressions at their root after we evaluate them
1. Is a [binary tree](https://en.wikipedia.org/wiki/Binary_tree) (convenient)
1. Is what everyone else uses

Our tree implementation is in [tree.go](./tree.go), along with some helper methods. To get from string expression to tree we will need a parser, which can be found in [parse.go](./parse.go) and is tested in [parse_test.go](./parse_test.go). Here is how the parser works:

```go
unparse(parse("fxyz")) // Returns `fxyz`
unparse(parse("(((fx)y)z)")) // Returns `fxyz`
```

## Section 3

> Now a sequence of particular functions of a very general nature will be introduced.

Up until now we have been talking about functions in the abstract. Schönfinkel now names and defines some functions we will be using in the rest of the paper. It is probably time we give these functions a name - "Combinators". You'll notice that Schönfinkel never uses this name, this term was popularized by Curry. Our type definition for a combinator can be found in [combinator.go](./combinator.go), which is repeated here:

```go
type Combinator struct {
    Name       string
    Arguments  []string
    Definition string
}
```

We can transform an expression involving our combinator into a reduced version using substitution via the following:

```go
var R = Combinator{
    Name:       "R", // `R` for "reverse"
    Arguments:  []string{"x", "y"}, // Takes two arguments
    Definition: "yz", // And swaps them
}
R.Transform("Rxy") // Returns `yx`
R.Transform("Ra(bc)") // Returns `bca`
```

### Identity

The first of which is the Identity function (or combinator), `I`. This function returns exactly what it is given, or

$$Ix = x$$

In our implementation we construct `I` like the following:

```go
var I = Combinator{
    Name:       "I",
    Arguments:  []string{"x"},
    Definition: "x",
}
```

and we can use it like this:

```go
I.Transform("II") // Returns `I`
```

### Constancy

Next up, the Constancy combinator. Schönfinkel calls this `C`, but in modern times it is refered to as `K`. For this guide and our implementation, I will use `K`. Taking two arguments, it always returns the first:

$$Kxy = x$$

Our implementation is:

```go
var K = Combinator{
    Name:       "K",
    Arguments:  []string{"x", "y"},
    Definition: "x",
}
```

### Interchange

Schönfinkel's explanation of the Interchange combinator, `T`, is quite hard to follow. It is simply:

$$Txyz = Txzy$$

`T` returns `x` (the first argument) with its inputs swapped (interchanged).

```go
var T = Combinator{
    Name:       "T",
    Arguments:  []string{"x", "y", "z"},
    Definition: "xzy",
}
```

### Composition

The Composition combinator, `Z`, is simply a means of shifting parentheses.

$$Zxyz = x(yz)$$

```go
var Z = Combinator{
    Name:       "Z",
    Arguments:  []string{"x", "y", "z"},
    Definition: "x(yz)",
}
```

### Fusion

Finally, the Fusion combinator, `S`, which Schönfinkel says will help us in the following way:

> Clearly, the practical use of the function S will be to enable us to reduce the number of occurrences of a variable - and to some extent also of a particular function - from several to a single one.

$$Sxyz = xz(yz)$$

```go
var S = Combinator{
    Name:       "S",
    Arguments:  []string{"x", "y", "z"},
    Definition: "xz(yz)",
}
```

### Transforming

So how does `Transform()` actually work? We parse our statement into a tree, and need an algorithm to do the actual work before we unparse back into a string. We can separate the work out into three separate bits:
1. Choosing a reduction strategy (which part of the expression to work on first)
1. Using the strategy, find reducible expressions (called redexes)
1. When found, rewrite the expression using the combinator definition

You can view the implementation in [reduce.go](./reduce.go) and [reduce_test.go](./reduce_test.go), but here is the algorithm spelled out:

#### Choosing a reduction strategy

There are three orthogonal choices to make when reducing an expression. They are as follows:

1. **Outer vs Inner**: Rewrite the outer expression first, or recurse through the expression tree first?
   - Example: $\underline{K(Ix)(Ix)}$ vs $\underline{K}(\underline{Ix})(\underline{Ix})$
2. **Left vs Right**: When recursing, choose the left or right subtree first?
   - Example: $(\underline{Ix})(Ix)$ vs $(Ix)(\underline{Ix})$
3. **Normal vs Applicative**: Reduce arguments before applying the function (combinator)?

Each combination of these choices should give us the same end result (via the [Church-Rosser theorum](https://en.wikipedia.org/wiki/Church%E2%80%93Rosser_theorem)), although some combinations may not be guaranteed to terminate.

For our implementation we will stick with *outermost leftmost normal-order* reduction.

#### Finding a redex

Each time we are ready to reduce a subtree, we first need to see if that particular subtree is capable of reducing (is it a **red**uceable **ex**pression)? To do this, we need to find the left-most leaf of the subtree, and see if it is our combinator. If it is, and there are enough arguments, we can perform the rewrite. Here are some examples:

Example 1 - The `K` combinator ($Kxy = x$) needs only 2 arguments, so the following tree can be reduced
```
Kabc
----
   /\
  /\ c (3'rd arg)
 /\ b (2'nd arg)
K  a (1'st arg)
```

Example 1 - The following `K` combinator does not have enough arguments and cannot be reduced
```
Ka
--
 /\
K  a (1'st arg)
```

#### Perform the rewrite

Finally we can actually rewrite our tree using the combinator definition. To follow our Example 1 above:

```
Kabc             ac
----             --
   /\             /\
  /\ c    ->     a  c
 /\ b
K  a
```

## Section 4

### Basis

So Schönfinkel has defined 5 combinators, but tells us in this section that we can actually derive some from the others (like the Sheffer stroke). Schönfinkel says that we actually only need `S` and `K`! In modern times we have actually not been able to do much better than just `S` and `K`. This is still the most common basis used in modern Combinatory Logic ([SKI combinator calculus](https://en.wikipedia.org/wiki/SKI_combinator_calculus)).

This brings us to the concept of a **basis**, which is the set of combinators one wants to use in their expressions. There are a few common bases:
- [**SKI**](https://en.wikipedia.org/wiki/SKI_combinator_calculus) - `S`, `K`, and `I`
- [**SK**](https://en.wikipedia.org/wiki/Combinatory_logic#Completeness_of_the_S-K_basis) - Just `S` and `K`
- [**BCKW**](https://en.wikipedia.org/wiki/B,_C,_K,_W_system) - Used by Curry, uses combinators `B` (Schönfinkel's `Z`), `C` (Schönfinkel's `T`), `K`, and `W` ($Wxy = xyy$)
- [**Iota**](https://en.wikipedia.org/wiki/Iota_and_Jot) - Just one combinator, $ιx = xSK$. More discussion on this in [section 6](./combinator.go)

You will see our implementaton of these commen bases in [combinator.go](./combinator.go) and [combinator_test.go](./combinator_test.go). You can use them as follows:

```go
// Create your own basis
schonfinkel := Basis{I, K, T, Z, S}
schonfinkel.Transform("S(ZZS)(KK)xyz") // Returns `xzy`

// Use a predefined basis
SK.Transform("SKKx") // Returns `x`
```

### The SK Basis

Schönfinkel now derives `I`, `Z`, and `T` in terms of `S` and `K`. I'll repeat them here:

- $I = SKK$
- $Z = S(KS)K$
- $T = S(ZZS)(KK)$ (using the definition for `Z` above)

These can be found in `TestSchonfinkel` in [combinator_test.go](./combinator_test.go)

## Section 5

This section is by far the most difficult in the paper. It took some time digesting this, alongside reading the ['1.2 The operators "nextand" and "U"'](https://plato.stanford.edu/entries/logic-combinatory/#OperNextU) section of [Stanford's Plato Combinatory Logic article](https://plato.stanford.edu/entries/logic-combinatory/).

Schönfinkel's aim in this section is twofold:
1. To be able to describe any first or second order logic formula in terms of just `K`, `S`, and `U`.
1. To eliminate bound variables (like $x$ below)

Of course, `U` is our universally quantified Sheffer stroke from section 1:

$$f(x) |^x g(x)$$

Which he defines in his new calculus as:

$$Ufg = fx |^x gx$$

The definition of this combinator is obviously different from the ones described in the previous section, mainly:

- The combinator definition is not in the form of a tree rewrite
- A bound variable $x$ has snuck its way into the combinator definition

So it seems that Schönfinkel is more or less going to "encode" first-order logic into his own logic system. This definitely does have merit (no bound variables, etc.), but I do get a feeling that he is just packaging them away into definitions. 

Lets give a simple example to understand what Schönfinkel is getting at:

| | |
|-|-|
| English | "For every number, that number can't be both even and odd" |
| First-order logic                | $(x)\overline{E(x) \\& O(x)}$                          |
| First-order logic (only Sheffer) | $Ex\|^xOx$ |
| Combinatory Logic                | $UEO$ |

So the bound $x$ is eliminated, and we have an equation with only `U`, and the other functions in the domain we are dealing with. `E` and `O` hold information about their domain of discourse, so no need to keep the $x$ around.

Schönfinkel now provides us with the methodology of how to move the the arguments of `U` around in particular cases using our existing combinators `K` and `S` (and those combinators that can be derived from `K` and `S`).

As far as our implementation goes, we can treat `U` as a constant as it has no meaning that is compatible with our tree definition. Maybe some day I will come back and implement the conversion of first-order logic to combinatory logic and vice versa.

The modern-day interpretation of this section is probably that combinatory logic is useful as an encoding mechanism to get as little syntax overhead as possible (and first-order logic is one such example). I find systems that build on top of the system (like the [Church encoding](https://en.wikipedia.org/wiki/Church_encoding)) much more interesting. Church encoding can be found in [combinator.go](./church.go) and [combinator_test.go](./church_test.go).

## Section 6

### `J` and Iota

In the final section, we see Schönfinkel attempt to reduce our combinator basis to just one combinator - `J`. `J` unfortunately has a definition that is incompatible with the previous explanation of combinators - `J` is made up of multiple definitions depending on its argument. This is an additional notion that is not quite in the spirit of the rest of the paper.

In modern times we have [iota](https://en.wikipedia.org/wiki/Iota_and_Jot), which very much accomplishes what Schönfinkel was trying to get at. Iota is the combinator `ι` (which for the remainder of this guide and implementation will be just `i` for ease of use). The iota combinator is defined as follows:

$$ix = xSK$$

so that

$$ii = I$$
$$i(i(ii)) = K$$
$$i(i(i(ii))) = S$$

I will note that Iota is technically "improper" (see [here](https://news.ycombinator.com/item?id=25335175) for discussion) as it is defined from other combinators. It is still quite amazing that we have a language with one character that can encompass all of computation.

### Further reduction

The final three paragraphs of the paper are not from Schönfinkel himself but written by the editor, [Heinrich Behmann](https://en.wikipedia.org/wiki/Heinrich_Behmann).

In the first paragraph, he articulates the fact that we can move all occurrences of `U` to the end of the expression, leave it off, and assume that our expression takes one argument, `U`.

The second paragraph aims to rid expressions of parentheses. One would want to rid parentheses for a few reasons:
1. So that any expression can be represented by a single number, usually with some [Gödel numbering](https://en.wikipedia.org/wiki/G%C3%B6del_numbering) scheme.
1. So that every possible expression is a valid expression.

Behmann suggests that we could simply leave out the parentheses, and let the combinators handle the work. Quine in his introduction had this to say:

> However, as Behmann recognized later in a letter to H. B. Curry, there is a fallacy here; the routine can generate new parentheses and not terminate

The last paragraph is not super interperable to me. If you have a good grasp of what this is saying please reach out!

### Iota, Jot, etc.

I think a good place to leave off is an exposition to [Chris Barker](https://en.wikipedia.org/wiki/Chris_Barker_(linguist))'s languages Iota, Jot, etc.

In the Iota language, `1` represents the iota combinator, and `0` represents the Polish operator. This Polish operator (described in the introduction by Quine) isn't actually representable as a combinator, and is a construct that exists outside of our language.

The only problem with Iota is that not every string of `1`'s and `0`'s is valid, so Barker created the language [Jot](https://esolangs.org/wiki/Jot), which is extremely similar.

## Overall Thoughts

1. Sometimes you can reduce logic down so far towards its fundamentals that you are actually going in the other direction of simplicity.
1. Reducing systems down to their simplest or most fundamental parts can be extremely valuable. As Wolfram pointed out, the [Sheffer stroke](https://en.wikipedia.org/wiki/Sheffer_stroke) (the [NAND gate](https://en.wikipedia.org/wiki/NAND_gate)) is all thats required to rederive all other logical connectives. This makes it the perfect gate to use as the core of nearly all transistors today.
   1. In general [Wolfram's "Combinators: A 100-Year Celebration"](https://www.youtube.com/watch?v=PG2G5xSz0NQ) has many extremely valuable insights on the application of Combinatory Logic.
1. Through this project I learned about [abstract rewriting systems](https://en.wikipedia.org/wiki/Abstract_rewriting_system), which are super interesting. Learning about rewriting systems definitely helped solidify my mental model of what combinators really are.
