package main

var test = `
one
two two
three three three
four four four four
abc
abd
abe
a
b
c
`

var test1 = `In this book we are going to read about some of the other people who
live in our own great country--Africa.  Africa is very, very large, so
big that no one would be able to go to all the places in it.  But
different people have been to different parts, and have told what they
saw where they went.  Wherever our home in Africa may be, if we walked
towards the sunrise--that is, towards the east--day after day, at last
we should reach the great salt sea.  Again, if we walked towards the
sunset in the west, we should at last get to the sea.  To the north,
again, is the sea, and to the south, the sea.  Whichever way we
walked, at last, after many months, we should be stopped by the sea.
But on our journey we should have met many different kinds of people,
and have seen many different customs.  In some places there would be
rivers, in some mountains, in some deserts, with no trees or grass to
be seen.  In these, people must make their homes in many ways, and
have many kinds of food and clothes.  Because we live in Africa, we
want to know about Africa and the people in it.  They are men and
women and children like ourselves, though the colour of their skins
may be lighter or darker than ours, and their languages quite
different.  But they, too, build houses and eat food and wear some
kind of dress, and it is interesting to know about their customs.  So
in this book we shall read about some of them and of how they live;
and, to help us to understand, we shall find with each part a picture
of the people we are reading about.  All the time we must remember
that we could get to see them for ourselves if we were strong enough
to walk so far, because they are all our own brothers and sisters in
Africa.`

var book = `
The Project Gutenberg EBook of People of Africa, by Edith A. How

Copyright laws are changing all over the world. Be sure to check the
copyright laws for your country before downloading or redistributing
this or any other Project Gutenberg eBook.

This header should be the first thing seen when viewing this Project
Gutenberg file.  Please do not remove it.  Do not change or edit the
header without written permission.

Please read the "legal small print," and other information about the
eBook and Project Gutenberg at the bottom of this file.  Included is
important information about your specific rights and restrictions in
how the file may be used.  You can also find out about how to make a
donation to Project Gutenberg, and how to get involved.


**Welcome To The World of Free Plain Vanilla Electronic Texts**

**eBooks Readable By Both Humans and By Computers, Since 1971**

*****These eBooks Were Prepared By Thousands of Volunteers!*****


Title: People of Africa

Author: Edith A. How

Release Date: October, 2004  [EBook #6693]
[Yes, we are more than one year ahead of schedule]
[This file was first posted on January 14, 2003]

Edition: 10

Language: English


*** START OF THE PROJECT GUTENBERG EBOOK, PEOPLE OF AFRICA ***




Produced by John Walker.



Line #1. . .Text begins on Line #228
            Production notes at line #16
            Explanation of typographical conventions at line #34

This electronic edition of  Edith  A. How's  _People  of  Africa_  was
produced  by John Walker in January 2003.  It follows the 1921 edition
(the only one of which I am aware) published in London by the  Society
for  Promoting  Christian  Knowledge  and in New York by the Macmillan
Company.  I have corrected two typographical errors  in  the  original
text: "sandstorm" was misspelled as "standstorm" on page 21 (section 1
of chapter III), and "bought" appeared where "brought" was intended on
page 33 (paragraph 3 of section 2 of chapter IV).

----------------------------------------------------------------------

                          PEOPLE OF AFRICA
                        Etext Production Notes

This  public  domain  Etext edition of Edith A. How's People of Africa
was prepared by:

        John Walker
        http://www.fourmilab.ch/

If  you discover any errors in this Etext, please report them to me by
E-mail.  If you're reporting a discrepancy between  the  Etext  and  a
modern  edition,  please  include  a complete citation of your source.
Upon  close  examination,  most  editions  contain  minor  errors  and
discrepancies which I've tried to correct in this Etext.  These Etexts
are part of the intellectual heritage we share as humans--please  help
to make them _perfectly embody_ the authors' legacies to the thousands
of generations and billions of readers whose lives they will enrich.

                      Beautifully Typeset Etexts
                      --------------------------

Free Plain Vanilla Etexts don't have to be austere and typographically
uninviting.  Most literature (as opposed to  scientific  publications,
for   example),   is   typographically  simple  and  can  be  rendered
beautifully into  type  without  encoding  it  into  proprietary  word
processor file formats or impenetrable markup languages.

This Etext is encoded in a form which  permits  it  to  be  both  read
directly   (Plain   Vanilla)   and   typeset   in   a  form  virtually
indistinguishable from printed editions of the work.

To create "typographically friendly" Etexts, I adhere to the following
rules:

    1.  Characters  follow the 8-bit ISO 8859/1 Latin-1 character set.
        ASCII is a proper subset of this character set, so any  "Plain
        ASCII"  file meets ths criterion by definition.  The extension
        to ISO 8859/1 is required so that  Etexts  which  include  the
        accented  characters  used  by  Western European languages may
        continue to be "readable by both humans and computers".

    2.  No  white  space  characters  other  than  blanks   and   line
        separators  are  used  (in  particular,  tabs  are expanded to
        spaces).

    3.  The text bracket sequence:
  <><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><>
        appears  both  before  and after the actual body of the Etext.
        This allows including an arbitrary prologue  and  epilogue  to
        the body of the document.

    4.  Normal body text begins in column 1 and is  set  ragged  right
        with  a  line  length  of  70  characters.   The  choice of 70
        characters is arbitrary and was made to avoid overly long  and
        therefore less readable lines in the Plain Vanilla text.

    5.  Paragraphs are separated by blank lines.

    6.  Centring,  right,  and  left  justification  is  indicated  by
        actually so-justifying the text within the 70 character  line.
        Left  justified  lines  should  start  in  column  2  to avoid
        confusion with paragraph body text.

    7.  Block  quotations  are  indented to start  in column 5 and set
        ragged right with a line length of 60 characters.

    8.  Preformatted tables begin with a line which starts in column 3
        and  contains at  least one  sequence of three  or more spaces
        between nonblanks.   The table is formatted verbatim until the
        next blank line.

    9.  Text set in italics is  bracketed  by  underscore  characters,
        "_".  These must match.

   10.  Footnotes  are  included  in-line,  bracketed  by  "[]".   The
        footnote appears at the point in the copy where  the  footnote
        mark  appears in the source text.  Footnotes may not be nested
        and may consist of only a single paragraph.

   11.  The  title  is  defined  as the sequence of lines which appear
        between the first text bracket "<><><>..." and a centred  line
        consisting exclusively of more than two equal signs "====".

   12.  The  author's name is the text which follows the line of equal
        signs marking the end of the  title  and  precedes  the  first
        chapter mark.  This may be multiple lines.

   13.  Chapters  are  delimited  by  a three line sequence of centred
        lines:
                           <Chapter number>
                         --------------------
                            <Chapter name>
        The line of minus signs must be centred and contain  three  or
        more  minus  signs  and  no  other characters apart from white
        space.  Chapter "numbers" need not be numeric--they can be any
        text.

   14.  Dashes in the text are indicated  in  the  normal  typewritten
        text  convention  of "--".  No hyphenation of words at the end
        of lines is done.

   15.  Ellipses are indicated by "..."; sentence-ending  ellipses  by
        "....".

   16.  Greek  letters  and  mathematical  symbols are enclosed in the
        brackets "\(" and "\)" and are expressed as their character or
        symbol  names in the LaTeX typesetting language.  For example,
        write the Greek word for "word" as:

                \( \lambda \acute{o} \gamma o \varsigma \)

        and the formula for the roots of a quadratic equation as:

                \( x_{1,2} = \frac{-b \pm \sqrt{b^2 - 4ac}}{2a} \)

        I  acknowledge that this provision is controversial.  It is as
        distasteful to me as I suspect it is to you.  In its  defence,
        let   me  treat  the  Greek  letter  and  math  formula  cases
        separately.  Using LaTeX encoding for Greek letters is  purely
        a  stopgap  until  Unicode  comes  into  common  use on enough
        computers so that we can  use  it  for  Etexts  which  contain
        characters  not in the ASCII or ISO 8859/1 sets (which are the
        7- and 8-bit subsets of Unicode, respectively).  If an  author
        uses  a Greek word in the text, we have two ways to proceed in
        attempting to meet the condition:

             The etext, when displayed, is clearly readable, and
             does *not* contain characters other than those
             intended by the author of the work, although....

        The first approach is to transliterate into  Roman  characters
        according  to  a  standard  table  such  as that given in _The
        Chicago Manual of  Style_.   This  preserves  readability  and
        doesn't  require  funny  encoding, but in a sense violates the
        author's   "original   intent"--the    author    could    have
        transliterated  the  word in the first place but chose not to.
        By transliterating we're reversing the author's decision.  The
        second  approach,  encoding  in  LaTeX  or  some  other markup
        language, preserves the distinction that the author wrote  the
        word  in  Greek  and  maintains  readability since letters are
        called out by their English language names, for the most part.
        Of  course LaTeX helps us only for Greek (and a few characters
        from other languages).  If you're faced with Cyrillic, Arabic,
        Chinese,  Japanese,  or  other  languages written in non-Roman
        letters, the only option (absent Unicode) is to transliterate.

        I   suggest  that  encoding  mathematical  formulas  as  LaTeX
        achieves the goal of "readable by humans" on the  strength  of
        LaTeX   encoding   being   widely  used  in  the  physics  and
        mathematics communities when writing formulas  in  E-mail  and
        other  ASCII  media.   Just as one is free to to transliterate
        Greek in an Etext, one can use ASCII artwork formulas like:

                                          ---------
                                     +   /  2
                                  -b - \/  b  - 4ac
                        x     =  ------------------
                         1,2            2a

        This is probably  a  better  choice  for  occasional  formulas
        simple enough to write out this way.  But to produce Etexts of
        historic  scientific  publications  such  as  Einstein's  "Zur
        Elektrodynamik  bewegter Krper" (the special relativity paper
        published in _Annalen der Physik_ in 1905), trying  to  render
        dozens of complicated equations in ASCII is not only extremely
        tedious  but in all likelihood counter-productive; ambiguities
        in trying to express complex equations would make it difficult
        for a reader to determine precisely what Einstein wrote unless
        conventions just as complicated (and harder to learn) as those
        of  LaTeX  were  adopted  for ASCII expression of mathematics.
        Finally, the choice of LaTeX encoding is made not  only  based
        on  its  existing  widespread  use  but because the underlying
        software that defines it (TeX and LaTeX) are entirely  in  the
        public  domain,  available in source code form, implemented on
        most commonly-available computers, and frozen by their authors
        so  that,  unlike  many  commercial  products,  the  syntax is
        unlikely to change in the future and obsolete current texts.

   17.  Other punctuation in the text consists only of the characters:

           . , : ; ? !  ' ( ) { } " + = - / * @ # $ % & ~ ^ | < >

        In other words, the characters:

           _ [ ] \

        are never used except in the special senses defined above.

   18.  Quote marks may be rendered explicitly as open and close quote
        marks with the sequences single quotes' or double quotes''.
        As long as quotes are balanced within a paragraph,  the  ASCII
        quote  character  "' may be used.  Alternating occurrences of
        this character  will  be  typeset  as  open  and  close  quote
        characters.   The open/close quote state is reset at the start
        of each paragraph, limiting the scope of errors  to  a  single
        paragraph and permitting continuation quotes'' when multiple
        paragraphs are quoted.

A program to translate Etexts prepared in this format into:

    LaTeX (and thence to PostScript or PDF, if you wish)
    HTML for posting on the Web
 or Palm Reader format for handhelds

may be downloaded from:

    http://www.fourmilab.ch/etexts/etset/

The program is in the public domain and includes complete source code.

<><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><>

                            PEOPLE OF AFRICA
                 =====================================
                         by Edith A. How, B.A.

                Universities' Mission to Central Africa

                    With Six Coloured Illustrations

                                 LONDON
               Society for Promoting Christian Knowledge

                      New York: The Macmillan Co.

                                  1921

                            ---------------
                                PREFACE

It is hoped that this book and its companion volume dealing with
non-African peoples will be the beginning of a series of simple,
readable accounts for Africans of some of the various objects of
general interest in the world of to-day.  There are many such works
published for the use of English and American children.  But the
native African has a totally different experience of life, and much
that is taken for granted by a child of a Northern civilized land
needs explanation to one used to tropical uncivilized surroundings.
Again, the African knows the essential operations of everyday life in
their simplest form, whereas the European knows them disguised by an
elaborate industrial system. All this makes books written for English
children almost unintelligible to a member of a primitive race.
These two volumes are far from perfect, but it has been difficult to
know always how to select wisely from the mass of material at hand.
They will have served, however, a useful purpose if they form a basis
for adaptations into the various African vernaculars, and afford an
inspiration for other works of a similar nature.  Thanks are due to
Miss K. Nixon Smith, of the Universities Mission to Central Africa,
for her kindness in criticizing the MSS. from her long experience of
the African outlook.

                                                          EDITH A. HOW
_June_, 1920.

                                  I
                             -----------
                             INTRODUCTION

In this book we are going to read about some of the other people who
live in our own great country--Africa.  Africa is very, very large, so
big that no one would be able to go to all the places in it.  But
different people have been to different parts, and have told what they
saw where they went.  Wherever our home in Africa may be, if we walked
towards the sunrise--that is, towards the east--day after day, at last
we should reach the great salt sea.  Again, if we walked towards the
sunset in the west, we should at last get to the sea.  To the north,
again, is the sea, and to the south, the sea.  Whichever way we
walked, at last, after many months, we should be stopped by the sea.
But on our journey we should have met many different kinds of people,
and have seen many different customs.  In some places there would be
rivers, in some mountains, in some deserts, with no trees or grass to
be seen.  In these, people must make their homes in many ways, and
have many kinds of food and clothes.  Because we live in Africa, we
want to know about Africa and the people in it.  They are men and
women and children like ourselves, though the colour of their skins
may be lighter or darker than ours, and their languages quite
different.  But they, too, build houses and eat food and wear some
kind of dress, and it is interesting to know about their customs.  So
in this book we shall read about some of them and of how they live;
and, to help us to understand, we shall find with each part a picture
of the people we are reading about.  All the time we must remember
that we could get to see them for ourselves if we were strong enough
to walk so far, because they are all our own brothers and sisters in
Africa.

Long ago most African peoples were shut off from the other people of
the world by the sea and the great sandy desert.  Only the people of
Egypt could meet and learn from the people of Europe and Asia.  So
while the Egyptians grew wise and clever, all the other Africans,
south of the desert, knew nothing except what they had learnt by
themselves.  Then Arabs began to cross the desert to get gold and
slaves from the dark-skinned Africans.  These Arabs taught them a
little.  But, later still, Europeans began to come in great ships over
the sea.  These came at first like the Arabs to trade, and afterwards
began to plant great fields of cotton and tobacco, which will not grow
in their own lands.  But they found the dark-skinned Africans were
still ignorant, and afraid of people of other races.  They were always
fighting among themselves, and no one could settle among them until
there was peace and safety.  At last the European nations made
agreements with the chiefs, so that now in nearly every part of Africa
there is a European governor to prevent wars and fighting.  Thus in
North Africa the governors are sent by France, in the Congo lands by
Belgium, in East Africa by England, in some other parts by Portugal.
These are different European nations who send men to keep peace, and
to make it possible to carry on trade.  Of course, the coming of the
Europeans has made great changes in the lives of the Africans.  In the
old times all the men were busy fighting, and often whole villages of
people were killed or made slaves.  Now there is no fighting, but
there is more need to work than before.  There are more people, and
less land for each family.  Europeans want workmen to help on their
great fields.  The Africans want many things now, which they did not
know about before, and they must have money to buy them.  So work for
money has taken the place of fighting.  Again, in some ways the
Europeans, enforcing peace and making many quick ways of travel, such
as good roads and bridges, have helped to weaken the power of the
chiefs.  Nobody likes changes to come, and the old people are always
sorry when their children begin new customs; but on the whole it is
good for Africans that other nations came to their country, because
they have brought peace in the place of war, and safety and freedom
instead of the old fear of death or slavery.

                                  II
                             -----------
                                EGYPT

                     1. The Country and its River

Egypt is a country in the north of Africa.  It has sea to the north
and sea to the east.  On the north it is called the Mediterranean Sea,
and on the east the Red Sea.  On the west is the great sandy desert
called the Sahara, and to the south are great forests and mountains.
Egypt itself is the land of the great River Nile.  There is very
seldom any rain there, and everyone has to get water from the great
river.  So all the people live near the Nile or the canals which lead
out of it.  A "canal" is a waterway, the channel of which has been dug
by men.  The big towns are where the river flows out into the sea, or
where a canal meets the main stream, because the people bring their
merchandise to market in boats.  All over the land are little
villages, where many people live and work in the fields to grow food.
Year by year when there is heavy rain in the mountains far away south,
the River Nile rises and floods the fields.  Then the people plant
their seed quickly and get a good harvest.  It is not difficult to
understand why the Egyptians love their great river, which gives them
water for their fields and carries them in their boats from place to
place.

                         2. Its Past History

Egypt is the only part of Africa that could be reached easily by
people in Europe and Asia, because in Egypt is the only place where
men could walk from Asia and Europe into Africa.  Even if they did not
want to walk, the sea was not too wide to cross in small boats.  In
the Bible we read how Abraham, who lived in Asia, walked to Egypt, and
later how Moses led the Children of Israel back to Asia.  Since that
time Europeans have cut a waterway for ships through this narrow neck
of land, which is called the Suez Canal.  So now people can no longer
walk from Asia to Africa, but in the old days the Egyptians grew wiser
than others in Africa, because they were more able to meet men from
other lands in Asia and Europe, and to learn something from them all.
So hundreds and hundreds of years ago there were people living in this
country of the Nile who were wise and great.  They built large cities
and temples and houses.  They knew how to write, and covered the walls
of their houses with writing.  Their letters were not like ours, but
were pictures of the things they were writing about.  They also built
huge stone tombs for their kings to be buried in, and these were
called "pyramids." The kings of Egypt were called "Pharaohs." When the
old Egyptians wrote books, instead of paper they used the dried leaves
of a reed called "papyrus," which grows in the Nile.  Several leaves
were fastened together to make a book.  These old writings on reeds
and on the walls have been found after lying buried in the sand, which
has covered so much of old Egypt.  The hot sand has kept them dry, and
prevented them being destroyed during hundreds of years.  By reading
these writings we are able to find out how these people lived so long
ago.  They had also a wonderful way of taking the waters of the Nile
in ditches over the whole land.  There is hardly any rain in Egypt,
and this Nile water prevented the country becoming so sandy and dry.
In those days Egypt was well-known for its wonderful harvests and
stores of food.

But though these people were wise in many ways, yet they were proud
and cruel to their enemies.  In the Bible we read how they treated the
Children of Israel in the time of Moses.  Perhaps this was because
they did not know God our Father, but worshipped many gods, whose
pictures and images were like animals.  Many of the great temples they
built for these gods are still standing, and when we see pictures of
them, we wonder at the skill of these people who lived so long ago.
Egypt was one of the first great countries to become Christian, and
many of the old heathen temples were turned into churches.  But at
last the Arabs, who were Mohammedans, conquered Egypt, and forced most
of the people to become Mohammedans too.  But some remained faithful
in spite of all, and these to-day are called "Copts," from the old
name for Egypt.  For hundreds of years these Copts have lived in a
country ruled by Moslem Arabs, or Turks, who hated their religion, but
they have been true to Christ through all.

There are people of all lands living in the towns of Egypt in these
days, for there is a great deal of business to be done in them.  But
the people who work in the fields are the children of the old
Egyptians, though they have forgotten their old wisdom and are now
very ignorant.

                        3. The People of Egypt

The Egyptians are a race different both from the dark-skinned people
of Africa and from Europeans.  They have olive skins, very dark,
almond-shaped eyes, and dark, straight hair.  Most of the men shave
their heads, and wear a turban or tarboosh as a covering.  The women
fasten a veil below their eyes, which falls over the lower part of
their face.  Both the men and the women wear several loose garments,
which cover the whole body from the neck to the feet.  All except the
very poor wear shoes.

In the towns there are a great many people, some very rich and others
very poor.  Often a city looks very beautiful, because the houses are
built of white or light-coloured stone or brick.  But they are close
together, and the streets are very narrow and dirty, and so the poor
people are often ill.  The houses are built in "storeys," one room on
the top of another, with steps leading to the upper rooms.  Often
there is a courtyard in the middle of the house, so that all the rooms
can have windows and light.  One part of the house is separated from
the rest for the women to use.  This is called the "hareem," and no
man, except the master of the house, is allowed to go into it.  All
rich Mohammedans have a separate part of their house for the women.  A
poor woman in all countries has plenty of work to do, but a rich lady
in Egypt has many servants, or slaves, to do the work, and, as she is
kept shut up in the "hareem" from the time she is ten or eleven years
old, she can learn very little, except how to do beautiful needlework.
She cannot help her husband and her sons to be wise and good, because
she does not know enough about life and work outside the "hareem." So
the Egyptian ladies have little to do and little to think about all
the day while their husbands are away, and they are often very dull.
But the town-people love their children very much, and Egyptian
children are taught always to love, honour, and obey their father and
mother.  An Egyptian man may have four wives, but generally he has
only one.

Until a few years ago, all Egyptians who had enough money used to buy
slaves to do their work.  Slaves could be bought or sold, or married
or given away, as if they were things instead of people.  Masters
could illtreat or even kill their slaves and not be punished, because
it was only as if they had broken their water-jar in a temper, and
that was no one else's business.  Often slaves were happy if they had
good masters, but it is a bad custom to take away a person's freedom
and treat him as if he had no soul.  During the last few years many
Europeans have been helping the Egyptians to improve their country,
and one of the changes has been to do away gradually with slavery.  No
one is now allowed to buy a slave, and anyone born in slavery can
become free if he wishes to do so.  Instead of slaves, people now have
servants who receive wages for their work.  These are free to leave
their master if he does not treat them well.  Although slavery is
dying out of Egypt, there are other parts of North Africa where the
old bad customs still exist, though the great European nations try to
prevent the public markets for slaves being held.  People are happiest
in countries where there are no slaves and everyone is free to do the
work for which he is best fitted.

In Egyptian households where there is more than one wife there is
often quarrelling.  The wives of one man all live in one "hareem," and
cannot help being jealous if they see their husband likes one better
than another.  Then there is quarrelling and ill-will among them.  As
the children grow up there is a further cause for jealousy, because
the mothers of boys are more important than those who have only
girl-children.  Children cannot respect their mothers if they often
see them quarrelling and jealous.  Again, there is always a
possibility that a husband may divorce his wife.  He is not likely to
do so if she has a boy-baby, but until she has, her position as a wife
is not very secure.  These bad marriage customs lead to much
unhappiness, and prevent the women of Egypt from doing so much good as
the women of some other lands are able to do.  We must not, of course,
think that all Egyptian homes are unhappy; probably many poor women
are quite glad when their husband brings another wife to help with the
work.  But where servants do the work, there are only the pleasures of
the home to be shared, and then jealousy will be likely to come.

                           4. The Big Towns

If we went for a walk in the narrow streets of an Egyptian city or big
town, we should see on either side open shops, each with its owner
ready to sell his goods.  Many of the people of the towns have shops
or trades.  They sell jewellery, furniture, cloth, and everything that
is wanted in the house for cooking.  In the streets there are some men
carrying drinking-water for sale, because it is hot walking about and
people get thirsty.  Others will be selling sweet-stuff made of sugar,
which everyone likes.  Others wait about ready to write letters for
people who cannot write for themselves, and there are always many
beggars.  Great steamers from other countries--England, France, India,
Japan--bring merchandise to Alexandria and Port Said, the seaports of
Egypt, and so people from these countries have shops and offices in
those towns.  Then the goods are taken by boats or trains to the
capital, Cairo, where the Sultan lives, and to other large towns.  In
all these towns there are hundreds of people, so that a man can only
know those who live near him or work with him.  Most of them are
unknown to one another and are like strangers, although they all live
in one town and can all speak Arabic.

                       5. Life in the Villages

The country-people of Egypt are very poor, and have to work very hard
all the year round in their fields.  Their houses are built of bricks
dried in the sun, plastered together with mud, and the roof is made of
plaited palm leaf.  Inside there is only one room, which has a big
oven made of mud with a flat top on which the father and mother sleep.
The work in the fields is very hard, as the ground has to be made
fertile by digging canals and ditches all over it to bring the water
from the Nile, because, you remember, there is no rain in Egypt.  When
the Nile begins to fall, the water has to be raised in baskets
fastened to a wheel or pole, and thrown on the ground.  In order to
get enough money, the people plant another kind of seed as soon as one
harvest is gathered; first, perhaps, planting wheat, then millet, or
cotton, then maize.  So the country-people in Egypt are always working
hard from sunrise to sunset all the year in their fields, and their
little children have to learn to mind sheep, goats, or cattle, and to
help in other ways as soon as they can walk alone.

Other men work on the Nile, carrying people or goods up and down the
river in boats from place to place.  This, again, is hard work, but
the boatmen seem very happy and often sing as they pass along.  People
in the country villages are ignorant, and very few can read or write.
Sometimes when the harvest has been bad and food is dear and scarce,
the people get deeply into debt.  There is a great deal of illness and
disease, but there are very few doctors and nurses to help people to
get well.  So the life of an Egyptian peasant is a hard one--a great
deal of work and very little time to rest, or play, or learn.  But
everyone has something to make him happy, and, unless there is famine
or pestilence, these people have their wives and children and home,
just as people have in England and other countries.  The only person
who need be unhappy is the one who has no one to love.

So we have learnt a little about that part of Africa called Egypt--the
land of the Nile--and about the people who live in it.  We must
remember that all the other people who live on the North Coast of
Africa, in Tunis, Algeria, and Morocco, are something like the
Egyptians, also speaking Arabic, and different from the dark-skinned
people who live farther south where it is very hot.

                                 III
                             -----------
                  THE SAHARA, THE GREAT SANDY DESERT

                      1. What the Desert is Like

In the last chapter we were reading about Egypt, and we said that on
the West of Egypt lay the Great Desert.  Now a desert is a place where
for some reason no food will grow.  In some deserts the soil is too
bad; in some the ground is covered with salt; in others, like the
Sahara, there are no rivers.  In some places in the Sahara there is
water coming up through a crack in the rocks.  This water is called a
"spring," and wherever one is found, trees and grass and food will
grow.  Such a place is called an "oasis." In the big oases there are
villages and towns.  But the sun is so hot that before the water from
the spring has flowed very far it is dried up, and beyond that nothing
will grow.  So when we think of the Sahara we have to try and picture
to ourselves a very big country, full of hills and valleys, but with
no rivers or lakes.  It is a journey of many months to cross the
Sahara, and day after day there is nothing to see but sand--sand, not
flat, but in ridges of hills like great waves of the sea.  When people
are travelling across this desert, they get very tired of looking at
nothing but sand all day.  Then, at last, as the sun sets, they reach
an oasis where there is water and bananas and date-trees, and perhaps
houses and people.  Sometimes great winds blow in the desert and bring
a sandstorm.  Then the sand beats hard against everything.  If
travellers meet a sandstorm, they have to throw themselves face
downwards on the ground to keep the sand out of their eyes and mouth.
Very often people who live in the desert have bad eyes, and many are
blind because of the sandstorms.

                        2. How the Desert Came

Long, long ago, the Sahara was not quite so dry as it is now.  There
were rivers then, which have dried up since.  When there was water,
food would grow, and people could keep sheep and cattle.  In those
days there were several large cities there.  But when the water began
to dry up, the ground became sandy and nothing would grow.  Then,
whenever the wind blew, the sand was carried along and began to cover
up the houses and temples.  The people had moved away because their
food would not grow, and soon the sand completely covered the old
cities.  For a long time they were buried, until some Europeans went
to see what they could find out about the people who lived there long
ago.  Then they dug and dug in the sand, and found the old houses and
temples.  But digging in the desert is very hard work, because it is
very hot, and there is very little water and food.  Often, too, a
great wind arises and brings a sandstorm.  Then the sand drifts back
again to the places already cleared.

                 3. The Desert Peoples (_a_) Berbers

It is surprising to find that there are a great many people living in
this desert region of North Africa.  There are three kinds of people
there.  Firstly, there are the Berbers, who live always in a little
town or village on a big oasis, and grow their own food.  Secondly,
there are the Bedouin, who live in large wandering tribes.  These keep
sheep and goats and camels, and stay on a small oasis until their
herds have eaten all the grass on it, and then move on to another
place.  Thirdly, there are the Arab traders, whose business is to go
south of the desert to get ivory and gold, and to take these back to
Egypt and to the great cities north of the desert to sell.  All these
people speak Arabic and are Mohammedans.

The Berbers who live in the towns on the great oasis, where there is a
large spring of water, are a different race from the Arabs, the
Egyptians, or the dark-skinned people of farther south.  They are much
darker-skinned than the Egyptians and the Bedouin.  In the past many
different races of South Europe, as well as the Arabs, have conquered
them and intermarried with them, but they still remain a distinct
race, though their customs are like those of other Moslems.  They make
their houses of bricks dried in the sun, and build them so close
together that people can step from one roof to another across the
street.  The roofs are flat, so that they can sit or sleep on them at
night when it is very hot inside the house.  All round the outside of
the towns are brick walls with gates that are shut at night for fear
of robbers.

These people live very much like the town-people in Egypt, only they
are much poorer.  They can buy things from the traders in the caravans
which stop at their village for the night, but as they cannot grow or
make many things to give in exchange, most people have to be content
with the earthenware cooking-pots and the cloth they can make
themselves.  The women draw water and prepare the food and look after
the children.  Then they weave flax and wool into cloth.  Their dress
is something like that of the poor Egyptians.  The children have to
herd the sheep and goats, which at night sleep in the house with their
owners.  The men hoe the gardens and grow the millet and barley for
food, and the flax for cloth.  The chief food of these people is bread
made of millet-flour kneaded with milk and baked in a hole in the
ground.  The flour is ground between two stones placed one on the top
of the other, the upper one having one or two handles by which it can
be moved round.  The people in these small, crowded towns in the
middle of the desert must live very narrow lives, and they do not know
much about anything outside their own village.  Journeys in the desert
are very dangerous because of sandstorms and the difficulty of finding
the way where there are no roads, and more especially because of
robbers.  So people never go on journeys unless they can join a big
company with plenty of men ready to fight if the robbers attack them.

                 4. The Desert Peoples (_b_) Bedouin

The second kind of people who have their home in the desert are the
Bedouin.  These are Arabs who once lived in another desert in Arabia,
but long, long ago many of them came to live in the Sahara.  The
Bedouin live in tents made of poles with dark cloth of goats' hair or
camels' hair spread across them for walls and roof.  They travel in
large tribes, and put up their tents on a small oasis where there is
no town.  These people still live as Abraham, Isaac, and Jacob lived
long ago, before the Israelites built their towns.  On the oasis their
camels, horses, sheep, and goats can find water to drink and grass to
eat.  When all the food has been eaten they pack up the tents and
everything they have and put it on the backs of the animals.  Then the
men and women and children all mount camels and horses and donkeys,
and the whole tribe moves to another oasis.  These people drink
camels' milk and eat the dates and bananas and other fruit they find
where they pitch their tents.  They also bring these fruits to the
Berber towns, and exchange them for flour to make bread and for coffee
to drink.  Coffee is a berry which is first roasted, then, when water
is boiled and poured on to it, it makes a strong, brown liquid which
Arabs and Europeans like to drink.  The women weave camels' hair into
clothes and blankets, and goats' hair into tent-covers.  The Bedouin
men are always ready to fight with their guns and lances; sometimes
they are robbers, but most of them travel from place to place, only
fighting if others attack them.  There is always a chief in each tribe
of Bedouin, and in each village of the Berbers, but away in the desert
there are many bands of robbers who will not obey any law, and
everyone has to fight for himself against these people.  The Bedouin
love their animals, especially their camels and their horses.  It is
quite natural that they should do so, because often a man would die in
the desert if his horse or camel would not work well and carry him
faithfully until they reached water.  Sometimes when the people lose
their way in the pathless sand, the horses and camels can find it.

                 5. The Desert Peoples (_c_) Traders

The third kind of people who are found in the Sahara are the traders.
These, like the Bedouin, are Arabs, but often their homes are in some
town, either on the edge of the desert or in Egypt.  They travel from
the great North African towns and from Egypt, across the desert to the
rich countries south of it, where the dark-skinned people live.
There, south of the Sahara, they buy ivory and dyed goat-skins and
other things in exchange for cloth and beads, and return with their
merchandise to the northern towns again.  Many years ago they used to
capture slaves, but they cannot often do so now, because the Christian
Europeans try to stop trading in slaves.  The journeys of the traders
take many months, because often they have to go by a long road in
order to find water.  So they travel from oasis to oasis seeking shade
and water.  Sometimes they have to ride three or four days to reach
the next drinking-place.  Then they have to carry water for themselves
in goat-skins.  The camels can live for a few days without water,
though they get very weak.  For this reason, everyone who makes long
journeys in the Sahara has to ride on a camel.  A horse can travel
more quickly, but he, like a man, must have water every day.  So the
camel is sometimes called the "Ship of the Desert," because he, best
of all, can carry men across the waterless sand.  When traders travel
across the desert with their merchandise, they are very much afraid of
the desert robbers, who steal what they can from travellers.  So they
journey in large companies called "caravans," with a paid guide to
show them the best and the quickest way from oasis to oasis, and with
many men armed with guns and spears paid to ride along by the side of
the camels carrying the merchandise, and to fight if robbers come to
steal.  These Sahara robbers are very bad people, who fight, and steal
all they can get, and always kill everyone they can.  So everyone who
crosses the Sahara has to be ready to fight for his life as well as
his property.  The desert is so vast, and has so many hills and
hiding-places, that it is easy for the robbers to get away after they
have robbed a caravan.  Then, as silence once more falls on the place
of the struggle, the cries of the jackals and hyenas and vultures are
heard, as they come from miles away drawn by the smell of blood.
Swiftly they gather to feed on the bodies of the slain, and soon the
wind blows the sand smooth and clean, where a few hours before it was
trampled and stained with blood.  Perhaps only a few whitened bones
remain to show what has happened.

                        6. The North of Africa

So we have learned something about the people who live in the North of
Africa.  In Egypt, the land of the great River Nile, the people can
grow rich and prosperous.  They have time to learn, but, except the
Copts, many of whom are goldsmiths, they seem to have quite forgotten
how to make the beautiful things the old Egyptians made.  In the
desert, the Sahara, there is little water, and life is very hard.  All
day people must work to get enough for food and clothes.  It is a land
without a king and without laws, where each must fight for himself.
Yet these people, on their long journeys through the waterless waste,
have learned to be very brave and fearless and strong.  They are
patient, and endure great hardships without grumbling.  They love
music, and often sing as they ride over the silent sand.  In the
evening they gather round the fire to tell stories of what happened
long ago.  The people of North Africa are all Arabs or Egyptians or
Berbers, with olive complexions and smooth, dark hair as a rule.  Next
we shall read about the very dark-skinned races who live farther
south, in Central Africa, where the sun is much hotter.

                                  IV
                             -----------
                      UGANDA, AN AFRICAN KINGDOM

                          1. Central Africa

In the last chapter we read that the Arab merchants crossed the desert
to buy ivory and goat-skins from the people who lived farther south.
In these next two chapters we shall read about these people south of
the desert.  Their land lies in the very middle of Africa, and so is
called Central Africa.  It is a beautiful country, with many rivers
and great lakes and mountains.  Central and West Africa are also the
very hottest part of this continent.  Now when plants have a lot of
water and a lot of sun they grow very quickly, and so Central Africa,
with its hot sun and its great rivers and lakes, is a land of great
forests.  In these forests there are lions and leopards, elephants,
and deer; and ivory and skins, as well as gold, have for many years
been sold by the Central Africans to the traders from the desert.  On
the eastern side of this country there are more mountains, lakes, and
small rivers; on the western side there are great rivers, all of which
join one very large one called the Congo.  In this chapter we shall
read about some of the people who live on the eastern side on the
shores of the largest of all the lakes--the one called Victoria
Nyanza.  These people are called the Baganda, and their country is
Uganda.

                            2. The Baganda

The Baganda are dark-skinned Africans.  They all belong to one tribe
and speak one language, but all around them are other Africans
belonging to different tribes and speaking different languages.  About
sixty years ago, when the grandfathers of the men who are alive now
were still young, the first Europeans went to Uganda.  Until that time
the tribes in Central Africa had spent most of their time fighting one
another, killing many and making others slaves.  Some of these slaves
were sold to the Arabs to take away to Zanzibar and across the sea, or
to take across the desert to Egypt.  Some tribes were much stronger
than others, and some of these drove everyone else out of the country
they had chosen for themselves and made a kingdom of it.  One of these
strong tribes was the Baganda.  Others liked to wander from place to
place, but the Baganda chose to settle down on the shores of the great
Lake Victoria Nyanza, and to stay there always.

When Europeans went to Uganda they found the Baganda had a king to
whom they paid great honour.  The king had many officers under him.
Some of these were the chiefs of different parts of the kingdom.
Others had special work to do--one to hear all the lawsuits and to
settle disputes, another to command the army.  Others had to work in
the king's household, to wait on his wives and children, or to beat
the big drum to call the people when the king wanted them, or to take
care that no one entered the palace unless the king wished them to do
so.  But whatever their work was, all the chiefs and officers and
people honoured and obeyed the king, and, because in this way everyone
was ready to fight or to work for the king and the rest of the nation,
the Baganda were one of the strongest and wisest of all the African
peoples.

The old dress of these people was a cloth, not sewn, but simply
twisted tight round their body under their arms, and reaching nearly
to the ground.  Sometimes it was fastened also by a belt round the
waist.  The cloth is made from the bark of certain trees soaked in
water and beaten hard for many days until it is soft and thin and
strong like woven cloth.  Their houses were round and built of reeds,
with steep roofs which nearly reached to the ground.  The smaller
villages had only a few people in them, everyone in each village
being related to the rest.  But the Baganda also had big towns, the
biggest to-day being Mengo, where the king lives.  Here there were
people gathered together for the king's work, and many others brought
food and bark-cloth to market to sell.  The houses of the king and
the great chiefs were large and beautifully decorated with plaited
reeds.

The chief food of the Baganda is plantains or bananas, which are
peeled when unripe and wrapped in smoke-dried banana leaves.  These
packets are slowly cooked with very little water in earthenware
cooking-pots.  When the food is cooked it is pressed and beaten, and
then the leaves are opened out and make a plate.  Other things, such
as beans and vegetables and fish, are cooked in the same way, wrapped
in banana leaves and then eaten with the bananas.

Some of the Baganda fish in the lake, and when they go on journeys it
is often quicker to travel by boat on the lake.  Many Africans can
only make boats out of rough tree-trunks with the inside scooped out,
but the Baganda had learnt to build long, narrow boats with high
carved wooden ends.  These canoes shot through the water very swiftly,
as twenty or thirty men paddled together in each boat.  It is well
they learnt to travel quickly, because the lake is very wide and
distances are great.  Often there are sudden, violent storms, which
would overturn a clumsy boat.  The carving on the boats and the
beautiful reed-work on the chiefs' houses were different from the work
of other African tribes.  When people begin to try to make things
beautiful as well as useful it is a sign that one day they will become
wise and great.

                     3. Europeans Come to Uganda

In the old days the Baganda, like other African people, thought there
were spirits in all the rivers and lakes and trees and everywhere,
which could help or hurt men.  The chief spirit they feared and to
whom they offered sacrifice was the spirit of their lake, Victoria
Nyanza.  Their witch-doctors told the people when they thought this
spirit was pleased or angry.  These witch-doctors were often bad and
cruel, and really cared more about getting all the power they could
over the king and people than for anything else.  Sometimes they said
that people must be killed as a sacrifice to the Spirit of the Lake.

When Europeans first went to Uganda, a few went to trade, but most
went to teach the Baganda about the Christians' God.  Many boys went
to their school near Mengo and were taught.  But the witch-doctors
grew frightened and persuaded the king to drive away all the
Europeans, and to kill the Baganda who would not worship the Lake
Spirit because they were Christians.  Mutesa the king did this,
killing the Christian Baganda boys very cruelly by burning them to
death, and killing the European, Bishop Hannington, when he came.  But
in a few years there were more Christians than before, and now in
Uganda the king and nearly all the chiefs and people are Christians,
as well as many of the tribes living near them to whom the Baganda
have sent teachers.  All through the Christian African kingdom there
are schools and hospitals.  The Baganda were always strong, and now so
many are Christians they have stopped fighting the other tribes and
killing and making slaves, and instead they spend their time learning
to make useful and beautiful things, which make their homes happier
and more comfortable to live in.  They quickly learn all they can from
Europeans and Indians, and to-day, in Mengo and in the other large
towns of Uganda, there are trains and motor-cars and stores, while
steamers on the lake bring European and Indian things quickly from the
coast towns.  There are many Europeans and Indians living in Uganda,
and this is a good thing, because when many people of different races
meet, they learn from one another and so grow wiser.

                      4. Europeans help Africans

In this chapter we have read about one of the wisest tribes of the
dark-skinned African people.  The Arabs in the north came to Africa
long ago from their own home in Asia, and the Europeans in the south
came from their home in Europe.  Both these races had learnt by
themselves a great deal more than the African race has done.  This is
partly because their homes were not so hot, and so they had to think
hard to get enough food and to keep warm.  It is partly due, too, to
the way in which for hundreds of years the people of Europe and Asia
have been able to read and write, and have met and learnt from one
another.  The Africans never found out how to write, and so could only
learn from each other by listening, never by reading.  They were shut
off from the rest of the world until one hundred years ago, and all
they knew they had found out for themselves.  But among the Africans
some learnt more than others, and the Baganda are a tribe who used
their minds as well as their bodies in becoming strong.  So by
thinking and learning they grew wise as well as powerful, and now
Europeans and Indians have come to their country they are able to
learn all these other races can teach them, which is far more than any
one race could find out alone.

                                  V
                             -----------
                       THE PEOPLE OF THE CONGO

                        1. Towards the Sunset

In the last chapter we read about some of the people who lived in the
Eastern lands south of the desert.  They were among the wisest of the
dark-skinned African tribes.  In this chapter we shall read about some
of the people who live in the Western part of Central Africa.  If the
Baganda walked day after day towards the sunset, they would reach the
land of the great River Congo.  This is not a narrow strip of land
along one river, like Egypt, but a very large country with many great
rivers, but all of these at last pour their waters into one very large
one, which is called the Congo.  Then the Congo takes all the water
from the whole land to the great salt sea.  Like Uganda this country
is very hot, and so, because there is so much sun and so much water,
there are great forests.  In places where there are no trees the grass
and maize grow much higher than a man's head.  In the forests there
are wild beasts--lions, leopards, elephants, and hippopotami--as well
as deer which are good to eat.  Many of the people spend most of their
time hunting in the forests for food and skins.

                       2. The Different Tribes

The people of the Congo are all dark-skinned Africans of the same race
as the Baganda, except two tribes which are quite different.  These
other people are called the Pigmies, which means they are very small.
None of the Congo people have made a kingdom of their own like the
Baganda.  They belong to different tribes, each with its own customs
and language.  Most of them wear a piece of bark-cloth or the skin of
an animal for clothing, but some wear very little, and paint or tattoo
their bodies.  Their houses are built of reeds, some tribes covering
the reed-walls with a thick plaster of mud, others leaving them
unplastered.  The roofs of some are thatched with the long grass of
the country, others are made of plaited palm-leaf mats.  Each tribe
has its own way of making a house, but no one builds very big houses
or large villages.  None of the houses last more than three or four
years; but these people do not want their houses to stand for many
years, because they are not like the Baganda who chose a country and
stay there always.  The Congo tribes move their villages after a few
years and live somewhere else.  So villages are always shifting, and
nothing they make is wanted to last long.  Some weave mats and baskets
out of palm-leaves or reeds; others make pottery; others make
iron-headed spears and hoes for their fields, but only a few things
that can easily be carried are wanted to last.  When the village
moves, most of the things must be left behind.  So, until a tribe
decides to stay always in one place, it does not as a rule learn to
make many useful and beautiful things.

Again, often men of different tribes build their villages near one
another, but the people of the two villages keep quite separate.  Each
has its own chief and follows its own customs.  Several villages of
one tribe may all obey a great chief, but no tribe has a chief so
powerful as the king of Uganda.  The Congo tribes have not learnt
nearly so much as some other African peoples.  The customs of each
tribe depend partly on which district of this large country they live
in.  Those who live near the salt sea eat sea-fish, and get salt by
boiling the sea-water in their cooking-pots until the pot is quite
dry, and then the salt is left behind after the water has gone.  It
was clever of those people to find out they could get salt that way.
Others, who live near the great rivers, make canoes out of the
tree-trunks with the inside hollowed out.  In these they go out and
catch river-fish to eat.  Others live in a country good for goats, and
these keep large herds of goats.  Some make good earthenware cups and
pots, others carve wooden ones.  Some wear ornaments made of shells,
some of beads, some of berries, some of teeth; everyone uses the
things he can get most easily.  But each tribe follows its own
customs, and despises those of its neighbours.  They are afraid and
jealous of each other, and there is constant fighting between the
various groups of villages.

Some tribes want to be peaceful, and these plant their food, which is
maize or millet, or some other grain which can be ground into flour,
then made into porridge.  Others are hunters or fishermen, and chiefly
eat meat or fish.  Some live by fighting other tribes, and capturing
their food and slaves.  Some of these are called cannibals, which
means they eat the flesh of human beings.  People who do this are
despised by all other races in the world, as they are so ignorant that
they do not know that it is wrong to eat other men.  Many of the
people of the Congo are not cannibals, but there is always war and
fighting between the different tribes, and it is dangerous to travel
because so many are always watching to rob and kill strangers.  The
lions and other wild beasts are dangerous, but the bands of fighting
men are still more to be feared.  Everything is wild and unsafe, and
there is no law outside the village, so each one has to protect
himself.  Among the dark-skinned Central African people each village
has a chief who keeps order within it, and often a group of villages
of one tribe has a great chief.  There are old laws and customs of
each tribe, and if anyone breaks one of these and injures someone
else, the chief calls him and asks all about it, and punishes the man
who did the wrong.

                            3. The Pigmies

Now we will think about the other two tribes who live in this country,
but who are of quite a different race from the others.  These little
red and black Pigmy peoples do not have villages at all.  They are all
hunters, and each man wanders with his wife and children wherever he
chooses.  Then, near the village of some chief of another tribe, he
collects grass and sticks, and builds a little house which is too
small for an ordinary man to stand upright inside.  The Pigmy people
are not so dark-skinned as the other races of Central Africa, and they
are very small, not so high as an ordinary man's shoulder.  They live
by hunting with a bow and arrow.  The Pigmy man respects the chief
whose village he settles in, but he does not fight for him or serve
him as the other people do in his village.  When he chooses, he leaves
that village and goes somewhere else.  If the Pigmies want fruit or
anything the villagers have, they shoot an arrow into it.  Then,
later, when they come to fetch it, they leave a packet of meat in
payment, for these little people never steal.  Although they live
peaceably with the other races, they speak their own language, and
never have anything to do with other villagers, and they only marry
among their own people.  The Pigmy men wear a small strip of cloth,
and the women wear a bunch of leaves for their clothes.  Most people
of Central Africa like to be clean, and when there is enough water
they always wash and bathe, but the Pigmies hate water and are always
very dirty.  They have no cooking-pots, but roast the meat they have
got from hunting on a stick over a fire.  These Pigmy people have
learnt less than any other tribe in Africa, for they do not even know
that it is better to live in villages with others of their own race,
which is the beginning of learning most other things.

                        4. Many still Ignorant

So in this chapter we have read about some other people who live in
the very hottest part of Africa.  The Baganda are among the cleverest
Central Africans, and these Pigmies and the cannibal tribes are among
the most ignorant.  But the Congo lands are very large, and there are
many different peoples; they often move their villages, and because
they hate one another they fight whenever they get the chance.  So
these people are still very ignorant and miserable.  When they find
out that it is better to be peaceful and work to help each other, then
they will be able to grow wise and strong like the other Central
African people in Uganda, and like the dark-skinned people of South
Africa whom we shall read about in the next chapter.

                                  VI
                             -----------
                   THE MINE-WORKERS OF SOUTH AFRICA

                   1. The Cooler Land of the South

The Congo rivers and another great river called the Zambezi stretch
right across Africa from east to west.  North of this the country is
called Central Africa, about some of whose people we have been
reading.  South of it across the Zambezi lies South Africa.  East and
west of this land is the salt sea, on the east called the Indian
Ocean, on the west the Atlantic Ocean.  As we travel south the country
gets narrower and narrower, until the two great oceans meet at the
Cape of Good Hope.  Near the Congo and the Zambezi towards Central
Africa the sun is very hot, but as we journey southwards it gets
cooler.  When we reach the colder lands of the south we find that the
grass and maize do not grow so tall, and that there are no great
forests.  For long distances the land stretches as far as we can see,
covered with short grass, but there are no trees.  This kind of
country is called "veld" in South Africa.  There are some waterless
deserts here, too, but none so large as the Sahara in North Africa.
In other parts there are rivers, though some of them dry up in the
summer and only have water in the rainy season.  In South Africa, as
in Central Africa, it rains some months of the year and is dry the
others.

                          2. Black and White

In South Africa there are two races of people living side by side.
First, there are dark-skinned Africans like those of Uganda and the
Congo.  These belong to many tribes, each speaking its own language.
Secondly, there are many Europeans who, about three hundred years ago,
began to come across the great salt sea to live in South Africa.
Their own countries in Europe were too small for all the people in
them, but South Africa is so large that there was plenty of room.
These Europeans live in houses of brick or stone, and wear the same
kind of clothes which are worn by the people in Europe.  Their skins
are lighter-coloured than even those of the Egyptians and Arabs of
North Africa, and their hair is straight and often very fair.  There
are two chief European peoples in South Africa, the English and the
Dutch.  These speak different languages, but many of them can speak
both.  Europeans, as perhaps you know, are very clever at making
machines of iron to work for them.  They have made motor-cars to carry
them quickly along ordinary roads, and another machine called an
"engine" which draws many cars on its own road, which is made of two
iron rails.

Among the African people of South Africa there are many different
customs, but most people live in their own villages very much like
those of Central Africa.  Some tribes keep great herds of cattle,
which find plenty of food on the grassy plains of the "veldt." Many
have learned to copy European customs, especially those living near
the great European towns.  Some go long distances to work in these
towns, especially in places where gold or other valuable things are
found under the ground in the "mines." It is about these men who work
on the mines that we will read now.

                         3. Work in the Mines

When men first found gold in the ground it was near the surface, and
was not very difficult to get.  But when this had all been taken, they
had to dig deeper and deeper, until at last they found it easier to
cut out roads and rooms far down underneath the ground, and to look
for the gold among the earth and stones they found there.  Perhaps you
wonder how the miners get so deep down in the earth every day.  There
are no steps, but they get into a kind of cage called a "lift," which
slips down on a rope skip into a deep hole called a "shaft," to where
they want to work.  It is a wonderful machine, something like a
motor-car, only it goes down into the earth instead of along the top.
When the men get out of the skip down in the mine, there are many
different roads in it, and each man has to go to his own part to work.
When he reaches his place he has to drill holes in the rock for the
dynamite which breaks up the rock, and the loose stones are taken away
along the roads to the lift and then up to the top.  There it is
stamped with great hammers into dust, and washed, until the gold-dust
is separated from the rest.  There are thousands of men, both
underground and at the top, always at work at the mines.

Down in the mines it is always dark because the sunlight cannot get
down there, and so the people have to use lanterns.  In the larger
openings there are lamps fixed to the walls and ceilings lighted by
"electricity." Although it is dark below the ground, we must not
think it is cold.  On the contrary, it is very hot and difficult to
breathe, because there is no wind, so that the bad air does not get
cleared away.  It is hot and stuffy, like a house where people have
been sleeping all night with no windows open.  When people first made
mines, a great many died because of the bad air and because of fires,
but now they have machines which blow good air down into the ground,
and electric and other lamps which do not set fire to things easily,
and so there are not many people killed in the mines now.
Nevertheless, it is very hard and tiring work, and men are often ill
because of the dust which fills the air they breathe.  So the
Europeans to whom the mines belong pay for doctors and hospitals
where the sick can be cared for until they are well.

Many valuable minerals, besides gold, are found in South Africa, but
the chief mines are for gold, diamonds, and coal.  Diamonds are
beautiful stones, clear like water, which flash red, blue, and green
when they are turned about.  They are very hard, and are sometimes
used to cut glass.  But they are valuable because European and Indian
ladies will pay large prices for them, as they like to wear them as
ornaments.  Coal is a hard, black, shiny mineral used for burning.  It
makes better fires than wood, and burns much longer.  These
three--gold, diamonds, and coal--are the chief things found in mines
in South Africa.  But in other countries men find iron and silver and
copper (of which pennies are made), and tin and salt, and many other
useful things, in mines dug deep under the ground.

                        4. How the Miners Live

People often come from very long distances to work in the diamond
mines at Kimberley and in the gold mines at Johannesburg.  Sometimes
they walk, but in South Africa there are railways and trains to take
people to all the large towns, and a person can travel in one day by
train as far as he could walk in three or four days.

Very few people spend all their lives at the mines.  Most of the
workers come for six months or a year, because they want money for
clothes or food, as well as to buy cattle to pay the dowry for the
girls they wish to marry.  When they arrive at the mines, after their
long journey, their names are written in a book as miners, and they
are given places where they can live.  If the men are single they live
together in a large compound, which is a place enclosed by walls and
gates.  In these compounds there are houses where the men sleep, and
places where they can do their washing, and the European mine masters
provide people to clean these houses and to do the cooking.

If the workman has a wife he is given a house in a mine village,
called a "location." A location or a compound is like a village with a
great number of houses placed close together along straight roads.
The houses are sometimes built of stones or bricks, but more often of
corrugated iron.

In each location there are hundreds of people who have come to work at
the mines for a few months from different parts of South Africa.  They
are all strangers to each other and speak many different languages.
Most of them try to copy the dress of Europeans; but as European
clothes are very expensive to buy and soon wear out, the natives often
look ragged and dirty in them.

These native workers in the mines are supplied with food, such as
maize, corn, and meal; but there are shops in the locations and
compounds where they can buy other food, such as tea, coffee, sugar,
and bread, and where they can also get clothes and other European
things.

There are hospitals with doctors and nurses at all the mines to attend
to the sick and the injured.  There are also schools for the children
in the location.  It is difficult to teach in these schools because
the children speak different languages, and their parents only stay
for a short time.  But a great many do learn to read, write, to do
sums, and to sew.

The country near the mines is very often dry and dusty.  There are no
fields nor trees, unless planted by Europeans.

There are many laws regulating the life and work of the native miner;
for example, he must go to work every day unless the doctor says he is
too ill to do so.  At night every one must be in the location, unless
he be given a letter, which is called a "pass," from his master giving
the reason why he is not in the location.

                      5. Strict Laws for Miners

The reason for these laws is that all these people are far away from
their homes, and often no one can speak their language.  Their
relations and chiefs are far away and cannot help them, and so the
Government has to make laws to prevent bad people robbing and perhaps
killing them.  Wherever there is a great deal of money, there are
always thieves and bad people.  So the Europeans who own the mines and
pay the workmen make these laws to protect their workmen, until their
time on the mines is finished, and they can go home to their own
chiefs again.  There are police ready to see that everyone obeys the
laws, and if they find bad people or thieves they take them to a
police-court and lock them up.

In all the other chapters we have read about people living in their
own homes with their own relations.  But in this chapter we read about
Africans who leave their homes to work on the mines.  They work hard
and live a very different life from that lived in their village.  They
see many different people of other countries, hear many languages, and
find out many new things.  But no one wants to make his home there.
High wages are paid for hard work, but everything is strange and
different, and each one longs for his home.  So everyone is glad when
at last his work is done and his wages paid, and he is free to go back
to his own village and the people he loves.  We must remember that
South Africa is a very large country with a great many Africans in it.
Large numbers do go to work on the mines for a time, as we have been
reading, but we must not forget that all these men have their homes in
villages scattered all over that great country.  In these villages
there are chiefs and customs very much like those of Central Africa.
But the great difference between South Africa and Central Africa is
that in cool South Africa Europeans can make their homes, and so the
Africans there see many European customs which they copy.  Trains make
it easy to go from one part of the country to another, and no tribe is
allowed to fight.  Where there is no fighting, people have tried to
learn and to grow wise.  The dark-skinned races of South Africa are
learning to be good workmen, and some to be wise enough to be teachers
and even doctors to serve and help their own people to lead happier
and more useful lives.

                                 VII
                             -----------
                   THE GREAT FARMS OF SOUTH AFRICA

                        1. The Two White Races

In the last chapter we read about some of the dark-skinned Africans
who live in South Africa, but we said also that there are many
Europeans living there too.  These Europeans came from two nations in
Europe--the English and the Dutch.  Now in South Africa they live side
by side, doing the same work, and all obeying and helping the
Government of South Africa, which is European.  For many years these
two nations kept separate, but the wisest men in each saw that this
was bad, and they decided to make one strong nation.  When Europeans
go to live in another country, they take all their own customs with
them, and so in South Africa there are cities and houses exactly like
those in their old homes in Europe.  In the towns many people live
together, drawn there by their work.  Some work on mines or railways,
some have shops, some have to keep the town clean and healthy.  In all
European towns there are shops, because in Europe and in India and
China no one can make everything he needs for himself.  Each man
learns to make one thing well, and spends all the day making one kind
of thing.  Then he sells what he has made, and buys from other people
all the other food and clothes he needs.  A country where people work
and live in this way is called civilized.  It is a good way to live,
because people do their work better and have more time to think and
learn from others.  In another book we will read about civilized
countries and the town people of Europe and Asia.  In this chapter we
will read about the Europeans on the great farms of South Africa, who
live far away from the towns.  These people are mostly Dutch or, as
they are sometimes called, Boers, but some of the farmers are English.

                        2. What a Farm is Like

Now a farm is a large stretch of land which belongs to one man, who
uses it either to grow food in the ground, or else to raise large
herds of cattle, or horses, or sheep.  In a civilized country people
cannot grow their own food, because they are busy all day with some
other trade.  So some people make it their work to grow large
quantities of food, and sell all they do not need themselves.  Cattle
are kept for their milk, which all Europeans drink.  The flesh of
cattle and sheep is used for food.  The skins of cattle and horses are
dried and made into leather for shoes and harness.  Cattle and horses
are also used to draw heavy carts and ploughs, and for riding long
distances.  A plough is a machine used to break up the ground ready
for sowing seed.  It is quicker and better than a hoe.  Sheep are used
as meat, and are kept especially for their wool.  This is sheared or
cut off every year, and is washed and spun and then woven into cloth.
Woollen cloth is much warmer and stronger than cotton, and in cooler
countries where Europeans can live people always need warm clothes
some months in the year, because the sun is low down in the sky, not
overhead, and the air is cold.  It is quite easy to see how useful
cattle and horses and sheep are in South Africa, and why some people
work to rear large herds.

On other farms where food is grown, some plant wheat or maize for
people to eat; some plant food for cattle to eat.  But a great many
farms grow maize, as this grows better than other grains in South
Africa.  Some parts of this country have great plains or low rolling
hills covered with short grass as far as you can see.  This kind of
land is called the "veldt." In other places there are dry, dusty
plains.  Everywhere there are hills formed of great mounds of huge
stones.  These are called "kopjes." For many months in the year there
is no rain, and the country becomes dusty and the smaller rivers dry
up; then at last the rain comes and the rivers are filled up with
water, and the whole land is covered with grass and flowers.  If at
times the rain is very late in coming, often whole farms are ruined
because the crops wither, or the cattle die, for want of water.

                     3. The Farmer and his Family

We said that a farm always belongs to one man, called the farmer.
This man lives with his wife and children in a brick or stone house in
the middle of his land.  Sometimes, when his children grow up, the
sons marry and bring their wives to live in the father's house, while
the daughters go away to live with their husbands on other farms.  The
girls who do not marry still live at home with their father and
mother.  So there are often many people living together in one great
farmhouse.  Each man and woman will have their own room to sleep in,
and everyone will eat together in a big room, not used for sleeping.
In the evening they all sit together to talk about what has been done
during the day.  Outside, not far away, there are huts for the
Africans who work on the farm, and sheds for the cattle and horses and
the carts and ploughs.  The Africans who work on the farms are not
like those who work on the mines for a while and then go home.  The
farm-workers usually make their homes where they work, living there
with their wives and children.  They have as a rule no other village
or chief of their own.  Their wives work in the farmer's house.

All the Europeans have some work to do.  The men see that the
ploughing and sowing is done well, and, because the farm is large,
this takes a long time.  They have to look after the cattle and horses
and sheep, and to take care that their food and water are good and
that their sleeping sheds are clean.  If the cattle get ill, sometimes
a whole herd will die, and the farmer will lose a great deal of money.
The children watch the herds while they are grazing, and take care
they do not stray too far away.  The women have to see after the
household.  There are always African women servants to help, but there
is a great deal of work in a European house.  In every room there are
many chairs and tables which have to be moved when the room is swept.
On all the beds there are blankets and white cotton sheets.  A white
cloth is spread on the table when food is to be eaten.  Europeans wear
many clothes.  All these have to be washed whenever they are dirty,
and so one person will be kept busy all day washing and ironing if
there are many people living on a farm.

Then Europeans eat three or four times a day, and have many different
kinds of food.  Maize or wheat flour is made into bread or cakes.
Meat is either roasted or boiled, and is often eaten with green
vegetables.  Sometimes meat and vegetables are cut up into small
pieces and all boiled together for a long time.  Then it is called
soup, and is eaten with a spoon.  Milk from the cattle is used to
drink, and is also made into butter and cheese, which are hard, and
can be eaten with bread.  Europeans drink coffee like the Arabs, or
tea which is made from the leaves of another plant.  When mealtime
comes all the family come to the big room where a large table is
covered with a white cloth.  The food is brought in a large bowl or
dish, and the farmer or his wife puts some on a plate for each person.
Europeans use knives and forks and spoons in eating food.  The men and
women and children all sit together round the table.  On the farms as
a rule there is no wood or coal to make fires, so the sweepings of the
cattle-shed are made into cakes and dried in the sun.  This makes very
good fuel for fires.

                     4. How South Africa is Ruled

The Europeans on the farms do not see many other people, as the farms
are very large and are long distances apart.  Sometimes the men have
to go to town to sell their grain or cattle and to buy other things,
but they cannot leave their work very often.  The children are taught
to read and write at home, and sometimes when they are big enough they
are sent away to school in some town.  There they will live with
children from many other parts of South Africa, and will learn that
their farm is only a little part of a very big country.  Europeans are
Christians, and the children are taught that they must love and help
their country and other people always.  It is because European
children are taught to be ready to give up everything, even their
lives, to help their country to be good and great, that the Christian
European nations have grown as strong and wise as they are.  The
countries of Europe learnt about Christ many hundreds of years ago.

We said that South Africa was ruled by Europeans.  Their king is King
George who lives in England, but he does not rule or make laws by
himself.  In South Africa and in each of his other countries, King
George sends a Governor, because he himself is so far away.  Then the
people of South Africa choose someone in each district to go and help
the Governor to rule wisely.  When all these men from different parts
meet together it is called a Parliament.  This Council or Parliament
decides everything about ruling the country, and tells the Governor
what it is best to do for all the people in South Africa.

So in thinking of South Africa we have to think of a nation of people,
each doing one particular kind of work which is needed both by himself
and by everyone else.  Everyone's work is useful to the whole nation,
whether he works in a town, or on a farm, or on a railway.  The great
towns are where people sell what they have made and buy what else they
need.  The farm families live far away from one another, growing food
or wool for the nation.  But they, too, meet from time to time, and
they read newspapers about what is done in the great towns.  Then,
when the time comes to choose the men for the Parliament to help the
Governor, farmers and townsmen in each district say which man they
wish to go to it.  In this way everyone can help the nation by his
work, and everyone can help to keep peace and justice in the country
and to prevent bad people hurting the weaker ones.

                                 VIII
                             -----------
                              CONCLUSION

Now our book is finished, and we have read about some of the other
people who also live in our country of Africa.  There in the north are
the Bedouin and the traders, always moving from waterspring to
waterspring across the sand of the great Sahara, ever on the watch
against robbers.  Next there are the Egyptians living on the great
River Nile: some in towns with shops and trades; some very poor in the
villages, planting their seed when the river rises.  All these
Northern people are Mohammedans and the men marry several wives, and
the women are veiled and live apart.

Farther south it is very hot, and is a land of great lakes and rivers.
Here we read about the Baganda, the dark-skinned Africans who learned
to make a strong nation where all the people helped each other and
obeyed their king.  These are now Christian, and are quickly learning
other things from the Christian European nations who trade with them.
Then we read about the tribes farther west in the land of the River
Congo.  These people still move their villages from time to time, and
each man makes only what he needs in his own home.  There is often
fighting between the tribes, and many people are killed.  These Congo
people have learnt very little, and some eat the flesh of men and
women, and the little Pigmies do not even live in villages, but each
family by itself.

Farther south still is the great country of South Africa.  Here it is
not so hot, and Europeans have made their homes in it.  There are
Africans living in tribes and villages, but learning to be peaceful
and to help each other by their work.  Many of these at times go to
work in the mines to find useful things deep down in the ground.
There are also the Europeans: some in towns, some in farms, all
European and African bound together in the great nation of South
Africa, each doing his own part of the nation's work.

So that in this great land of Africa we have people living very
different kinds of life, in the deserts, in the forests of the Congo,
in Uganda and on the Nile, in the mines of South Africa, and on the
great farms on the veld and in the great towns.  The country itself is
different in different parts: the sand in the north; Central Africa,
with its hot sun and its lakes and rivers and mountains and forests;
South Africa, with its great grassy plains, and the mines and towns
joined by the railways which make it easy to get quickly to places far
away.  Yet, although the people of Africa have such different homes,
we must remember that they are very much like ourselves.  They wear
other clothes and speak other languages, but all love their families,
and each is doing his best to make his home a happy place in which he
can live.

                     Printed in Great Britain by
             Billing and Sons, Ltd., Guildford and Esher

<><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><><>




*** END OF THE PROJECT GUTENBERG EBOOK, PEOPLE OF AFRICA ***

This file should be named 6693.txt or 6693.zip

Project Gutenberg eBooks are often created from several printed
editions, all of which are confirmed as Public Domain in the US
unless a copyright notice is included.  Thus, we usually do not
keep eBooks in compliance with any particular paper edition.

We are now trying to release all our eBooks one year in advance
of the official release dates, leaving time for better editing.
Please be encouraged to tell us about any error or corrections,
even years after the official publication date.

Please note neither this listing nor its contents are final til
midnight of the last day of the month of any such announcement.
The official release date of all Project Gutenberg eBooks is at
Midnight, Central Time, of the last day of the stated month.  A
preliminary version may often be posted for suggestion, comment
and editing by those who wish to do so.

Most people start at our Web sites at:
http://gutenberg.net or
http://promo.net/pg

These Web sites include award-winning information about Project
Gutenberg, including how to donate, how to help produce our new
eBooks, and how to subscribe to our email newsletter (free!).


Those of you who want to download any eBook before announcement
can get to them as follows, and just download by date.  This is
also a good way to get them instantly upon announcement, as the
indexes our cataloguers produce obviously take a while after an
announcement goes out in the Project Gutenberg Newsletter.

http://www.ibiblio.org/gutenberg/etext04 or
ftp://ftp.ibiblio.org/pub/docs/books/gutenberg/etext04

Or /etext03, 02, 01, 00, 99, 98, 97, 96, 95, 94, 93, 92, 92, 91 or 90

Just search by the first five letters of the filename you want,
as it appears in our Newsletters.


Information about Project Gutenberg (one page)

We produce about two million dollars for each hour we work.  The
time it takes us, a rather conservative estimate, is fifty hours
to get any eBook selected, entered, proofread, edited, copyright
searched and analyzed, the copyright letters written, etc.   Our
projected audience is one hundred million readers.  If the value
per text is nominally estimated at one dollar then we produce $2
million dollars per hour in 2002 as we release over 100 new text
files per month:  1240 more eBooks in 2001 for a total of 4000+
We are already on our way to trying for 2000 more eBooks in 2002
If they reach just 1-2% of the world's population then the total
will reach over half a trillion eBooks given away by year's end.

The Goal of Project Gutenberg is to Give Away 1 Trillion eBooks!
This is ten thousand titles each to one hundred million readers,
which is only about 4% of the present number of computer users.

Here is the briefest record of our progress (* means estimated):

eBooks Year Month

    1  1971 July
   10  1991 January
  100  1994 January
 1000  1997 August
 1500  1998 October
 2000  1999 December
 2500  2000 December
 3000  2001 November
 4000  2001 October/November
 6000  2002 December*
 9000  2003 November*
10000  2004 January*


The Project Gutenberg Literary Archive Foundation has been created
to secure a future for Project Gutenberg into the next millennium.

We need your donations more than ever!

As of February, 2002, contributions are being solicited from people
and organizations in: Alabama, Alaska, Arkansas, Connecticut,
Delaware, District of Columbia, Florida, Georgia, Hawaii, Illinois,
Indiana, Iowa, Kansas, Kentucky, Louisiana, Maine, Massachusetts,
Michigan, Mississippi, Missouri, Montana, Nebraska, Nevada, New
Hampshire, New Jersey, New Mexico, New York, North Carolina, Ohio,
Oklahoma, Oregon, Pennsylvania, Rhode Island, South Carolina, South
Dakota, Tennessee, Texas, Utah, Vermont, Virginia, Washington, West
Virginia, Wisconsin, and Wyoming.

We have filed in all 50 states now, but these are the only ones
that have responded.

As the requirements for other states are met, additions to this list
will be made and fund raising will begin in the additional states.
Please feel free to ask to check the status of your state.

In answer to various questions we have received on this:

We are constantly working on finishing the paperwork to legally
request donations in all 50 states.  If your state is not listed and
you would like to know if we have added it since the list you have,
just ask.

While we cannot solicit donations from people in states where we are
not yet registered, we know of no prohibition against accepting
donations from donors in these states who approach us with an offer to
donate.

International donations are accepted, but we don't know ANYTHING about
how to make them tax-deductible, or even if they CAN be made
deductible, and don't have the staff to handle it even if there are
ways.

Donations by check or money order may be sent to:

Project Gutenberg Literary Archive Foundation
PMB 113
1739 University Ave.
Oxford, MS 38655-4109

Contact us if you want to arrange for a wire transfer or payment
method other than by check or money order.

The Project Gutenberg Literary Archive Foundation has been approved by
the US Internal Revenue Service as a 501(c)(3) organization with EIN
[Employee Identification Number] 64-622154.  Donations are
tax-deductible to the maximum extent permitted by law.  As fund-raising
requirements for other states are met, additions to this list will be
made and fund-raising will begin in the additional states.

We need your donations more than ever!

You can get up to date donation information online at:

http://www.gutenberg.net/donation.html


***

If you can't reach Project Gutenberg,
you can always email directly to:

Michael S. Hart <hart@pobox.com>

Prof. Hart will answer or forward your message.

We would prefer to send you information by email.


**The Legal Small Print**


(Three Pages)

***START**THE SMALL PRINT!**FOR PUBLIC DOMAIN EBOOKS**START***
Why is this "Small Print!" statement here? You know: lawyers.
They tell us you might sue us if there is something wrong with
your copy of this eBook, even if you got it for free from
someone other than us, and even if what's wrong is not our
fault. So, among other things, this "Small Print!" statement
disclaims most of our liability to you. It also tells you how
you may distribute copies of this eBook if you want to.

*BEFORE!* YOU USE OR READ THIS EBOOK
By using or reading any part of this PROJECT GUTENBERG-tm
eBook, you indicate that you understand, agree to and accept
this "Small Print!" statement. If you do not, you can receive
a refund of the money (if any) you paid for this eBook by
sending a request within 30 days of receiving it to the person
you got it from. If you received this eBook on a physical
medium (such as a disk), you must return it with your request.

ABOUT PROJECT GUTENBERG-TM EBOOKS
This PROJECT GUTENBERG-tm eBook, like most PROJECT GUTENBERG-tm eBooks,
is a "public domain" work distributed by Professor Michael S. Hart
through the Project Gutenberg Association (the "Project").
Among other things, this means that no one owns a United States copyright
on or for this work, so the Project (and you!) can copy and
distribute it in the United States without permission and
without paying copyright royalties. Special rules, set forth
below, apply if you wish to copy and distribute this eBook
under the "PROJECT GUTENBERG" trademark.

Please do not use the "PROJECT GUTENBERG" trademark to market
any commercial products without permission.

To create these eBooks, the Project expends considerable
efforts to identify, transcribe and proofread public domain
works. Despite these efforts, the Project's eBooks and any
medium they may be on may contain "Defects". Among other
things, Defects may take the form of incomplete, inaccurate or
corrupt data, transcription errors, a copyright or other
intellectual property infringement, a defective or damaged
disk or other eBook medium, a computer virus, or computer
codes that damage or cannot be read by your equipment.

LIMITED WARRANTY; DISCLAIMER OF DAMAGES
But for the "Right of Replacement or Refund" described below,
[1] Michael Hart and the Foundation (and any other party you may
receive this eBook from as a PROJECT GUTENBERG-tm eBook) disclaims
all liability to you for damages, costs and expenses, including
legal fees, and [2] YOU HAVE NO REMEDIES FOR NEGLIGENCE OR
UNDER STRICT LIABILITY, OR FOR BREACH OF WARRANTY OR CONTRACT,
INCLUDING BUT NOT LIMITED TO INDIRECT, CONSEQUENTIAL, PUNITIVE
OR INCIDENTAL DAMAGES, EVEN IF YOU GIVE NOTICE OF THE
POSSIBILITY OF SUCH DAMAGES.

If you discover a Defect in this eBook within 90 days of
receiving it, you can receive a refund of the money (if any)
you paid for it by sending an explanatory note within that
time to the person you received it from. If you received it
on a physical medium, you must return it with your note, and
such person may choose to alternatively give you a replacement
copy. If you received it electronically, such person may
choose to alternatively give you a second opportunity to
receive it electronically.

THIS EBOOK IS OTHERWISE PROVIDED TO YOU "AS-IS". NO OTHER
WARRANTIES OF ANY KIND, EXPRESS OR IMPLIED, ARE MADE TO YOU AS
TO THE EBOOK OR ANY MEDIUM IT MAY BE ON, INCLUDING BUT NOT
LIMITED TO WARRANTIES OF MERCHANTABILITY OR FITNESS FOR A
PARTICULAR PURPOSE.

Some states do not allow disclaimers of implied warranties or
the exclusion or limitation of consequential damages, so the
above disclaimers and exclusions may not apply to you, and you
may have other legal rights.

INDEMNITY
You will indemnify and hold Michael Hart, the Foundation,
and its trustees and agents, and any volunteers associated
with the production and distribution of Project Gutenberg-tm
texts harmless, from all liability, cost and expense, including
legal fees, that arise directly or indirectly from any of the
following that you do or cause:  [1] distribution of this eBook,
[2] alteration, modification, or addition to the eBook,
or [3] any Defect.

DISTRIBUTION UNDER "PROJECT GUTENBERG-tm"
You may distribute copies of this eBook electronically, or by
disk, book or any other medium if you either delete this
"Small Print!" and all other references to Project Gutenberg,
or:

[1]  Only give exact copies of it.  Among other things, this
     requires that you do not remove, alter or modify the
     eBook or this "small print!" statement.  You may however,
     if you wish, distribute this eBook in machine readable
     binary, compressed, mark-up, or proprietary form,
     including any form resulting from conversion by word
     processing or hypertext software, but only so long as
     *EITHER*:

     [*]  The eBook, when displayed, is clearly readable, and
          does *not* contain characters other than those
          intended by the author of the work, although tilde
          (~), asterisk (*) and underline (_) characters may
          be used to convey punctuation intended by the
          author, and additional characters may be used to
          indicate hypertext links; OR

     [*]  The eBook may be readily converted by the reader at
          no expense into plain ASCII, EBCDIC or equivalent
          form by the program that displays the eBook (as is
          the case, for instance, with most word processors);
          OR

     [*]  You provide, or agree to also provide on request at
          no additional cost, fee or expense, a copy of the
          eBook in its original plain ASCII form (or in EBCDIC
          or other equivalent proprietary form).

[2]  Honor the eBook refund and replacement provisions of this
     "Small Print!" statement.

[3]  Pay a trademark license fee to the Foundation of 20% of the
     gross profits you derive calculated using the method you
     already use to calculate your applicable taxes.  If you
     don't derive profits, no royalty is due.  Royalties are
     payable to "Project Gutenberg Literary Archive Foundation"
     the 60 days following each date you prepare (or were
     legally required to prepare) your annual (or equivalent
     periodic) tax return.  Please contact us beforehand to
     let us know your plans and to work out the details.

WHAT IF YOU *WANT* TO SEND MONEY EVEN IF YOU DON'T HAVE TO?
Project Gutenberg is dedicated to increasing the number of
public domain and licensed works that can be freely distributed
in machine readable form.

The Project gratefully accepts contributions of money, time,
public domain materials, or royalty free copyright licenses.
Money should be paid to the:
"Project Gutenberg Literary Archive Foundation."

If you are interested in contributing scanning equipment or
software or other items, please contact Michael Hart at:
hart@pobox.com

[Portions of this eBook's header and trailer may be reprinted only
when distributed free of all fees.  Copyright (C) 2001, 2002 by
Michael S. Hart.  Project Gutenberg is a TradeMark and may not be
used in any sales of Project Gutenberg eBooks or other materials be
they hardware or software or any other related product without
express permission.]

*END THE SMALL PRINT! FOR PUBLIC DOMAIN EBOOKS*Ver.02/11/02*END*
`
