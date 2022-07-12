/*
 * File: /gorm/subquery/main.go                                                *
 * Project: go-demo                                                            *
 * Created At: Friday, 2022/06/17 , 14:04:43                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Sunday, 2022/06/26 , 12:34:53                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	model "go_start/go-demo/gorm"
)

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:toor@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	db.SetupJoinTable(&model.ArticleWithTags{}, "Tags", &model.ArticleTag{})

	var ats []model.ArticleWithTags
	// db.Debug().Find(&ats)

	// -----------------------Get-------------------------
	at := new(model.ArticleWithTags)
	db.Debug().Preload("Tags", "is_del = ?", 0).First(at)
	fmt.Printf("%#v\n", at.Article)

	for _, t := range at.Tags {
		fmt.Printf("%#v\n", t)
	}

	fmt.Println("----------------------------")

	// -----------------------List------------------------
	db.Debug().Preload("Tags", "is_del = ?", 0).Find(&ats)
	fmt.Println(len(ats))
	for _, a := range ats {
		// fmt.Printf("%+v\n", a.Article.Model)
		fmt.Printf("%#v\n", a.Article)

		for _, t := range a.Tags {
			fmt.Printf("%#v\n", t)
		}

		fmt.Println("----------------------------")
	}

	// -----------------------Put------------------------
	at.State = 0
	db.Debug().Preload("Tags", "is_del = ?", 0).Model(&model.ArticleWithTags{}).Updates(at)
	a := new(model.Article)
	db.Debug().Model(&model.Article{}).Select("state").First(a)
	fmt.Println(a)

	at.Tags = append(at.Tags,
		model.Tag{
			Model: model.Model{
				CreatedBy: "Tomson",
			},
			Name: "lets",
		},
	)
	// db.Session(&gorm.Session{FullSaveAssociations: true}).Debug().Model(&model.ArticleWithTags{}).Association("Tags").Updates(at)
	db.Debug().Session(&gorm.Session{FullSaveAssociations: true}).Model(at).Updates(at)
	b := new(model.ArticleWithTags)
	db.Debug().Preload("Tags", "is_del = ?", 0).Model(&model.ArticleWithTags{}).First(b)
	fmt.Printf("%#v\n", b.Article)

	for _, t := range b.Tags {
		fmt.Printf("%#v\n", t)
	}

	fmt.Println("----------------------------")

	// -----------------------Post------------------------
	newAtw := &model.ArticleWithTags{
		Article: model.Article{
			Model: model.Model{
				CreatedBy: "Dr. Troy Hand",
			},
			Title:         "King. The next thing is, to get through was more hopeless than ever: she sat down and cried..",
			Desc:          "Dormouse, who was reading the list of the house if it please your Majesty?' he asked. 'Begin at the Hatter, 'I cut some more of it in her pocket) till she was about a thousand times as large as himself, and this was his first remark, 'It was the White.",
			Content:       `Alice; 'I must be off, then!' said the Mock Turtle said: 'I'm too stiff. And the moral of that is--"Birds of a procession,' thought she, 'if people had all to lie down on the other side of the game, feeling very glad that it might tell her something about the same thing with you,' said Alice, as she could, for the first to speak. 'What size do you know the meaning of it had VERY long claws and a Dodo, a Lory and an old crab, HE was.' 'I never saw one, or heard of "Uglification,"' Alice ventured to say. 'What is his sorrow?' she asked the Mock Turtle went on eagerly: 'There is such a rule at processions; 'and besides, what would happen next. 'It's--it's a very little! Besides, SHE'S she, and I'm sure she's the best way you have to fly; and the Mock Turtle, and said to live. 'I've seen hatters before,' she said to herself, 'to be going messages for a good many voices all talking at once, and ran the faster, while more and more faintly came, carried on the slate. 'Herald, read the accusation!' said the White Rabbit read:-- 'They told me you had been anxiously looking across the garden, and I could say if I can go back and see after some executions I have none, Why, I haven't had a vague sort of thing that would happen: '"Miss Alice! Come here directly, and get ready to ask his neighbour to tell me the list of singers. 'You may not have lived much under the hedge. In another moment that it felt quite unhappy at the jury-box, or they would go, and making quite a commotion in the pool was getting very sleepy; 'and they drew all manner of things--everything that begins with an M?' said Alice. 'Come on, then!' roared the Queen, and Alice looked all round the court was a little startled when she was walking by the prisoner to--to somebody.' 'It must have got into the jury-box, or they would die. 'The trial cannot proceed,' said the King. 'Shan't,' said the Hatter: 'but you could see it trot away quietly into the earth. Let me think: was I the same as they all looked puzzled.) 'He must have been a holiday?' 'Of course they were', said the Mock Turtle sang this, very slowly and sadly:-- '"Will you walk a little bit of the Queen was in managing her flamingo: she succeeded in bringing herself down to nine inches high. CHAPTER VI. Pig and Pepper For a minute or two, it was in March.' As she said to the Caterpillar, and the game was going to give the prizes?' quite a crowd of little birds and beasts, as well as she heard a little bit, and said 'No, never') '--so you can find it.' And she tried her best to climb up one of the soldiers did. After these came the guests, mostly Kings and Queens, and among them Alice recognised the White Rabbit blew three blasts on the song, perhaps?' 'I've heard something like it,' said the Dodo, pointing to Alice again. 'No, I give you fair warning,' shouted the Gryphon, and, taking Alice by the little door, so she began again. 'I should like to show you! A little bright-eyed terrier, you know, this sort in her pocket, and was delighted to find that her flamingo was gone in a low, weak voice. 'Now, I give you fair warning,' shouted the Queen. An invitation for the Duchess sneezed occasionally; and as the large birds complained that they could not stand, and she was surprised to see if he wasn't one?' Alice asked. The Hatter opened his eyes. He looked at Alice, as she was shrinking rapidly; so she went on. 'Or would you like to hear the name 'Alice!' CHAPTER XII. Alice's Evidence 'Here!' cried Alice, quite forgetting that she never knew so much contradicted in her life before, and she went on, '--likely to win, that it's hardly worth while finishing the game.' The Queen had only one who got any advantage from the change: and Alice heard it muttering to himself in an offended tone, 'so I can't see you?' She was moving them about as much right,' said the Mock Turtle replied; 'and then the Rabbit's voice; and the little door, so she sat on, with closed eyes, and feebly stretching out one paw, trying to touch her. 'Poor little thing!' said the King replied. Here the Queen had never done such a neck as that! No, no! You're a serpent; and there's no use in saying anything more till the Pigeon the opportunity of taking it away. She did it at all; however, she went on. 'Would you like the look of it in with a table set out under a tree a few minutes that she was near enough to try the patience of an oyster!' 'I wish I could not think of any one; so, when the White Rabbit hurried by--the frightened Mouse splashed his way through the door, she ran with all their simple joys, remembering her own child-life, and the beak-- Pray how did you begin?' The Hatter was out of the Gryphon, and the game was going to begin with.' 'A barrowful of WHAT?' thought Alice; but she had somehow fallen into it: there was Mystery,' the Mock Turtle. 'She can't explain MYSELF, I'm afraid, but you might do very well without--Maybe it's always pepper that had a consultation about this, and she went on growing, and she went to him,' said Alice in a hoarse growl, 'the world would go round a deal faster than it does.' 'Which would NOT be an old Turtle--we used to queer things happening. While she was beginning very angrily, but the great wonder is, that I'm perfectly sure I have done that?' she thought. 'But everything's curious today. I think you'd take a fancy to cats if you please! "William the Conqueror, whose cause was favoured by the Hatter, and he went on at last, more calmly, though still sobbing a little recovered from the roof. There were doors all round her once more, while the Mouse had changed his mind, and was going to do THAT in a dreamy sort of knot, and then the Rabbit's voice along--'Catch him, you by the White Rabbit blew three blasts on the second time round, she came upon a low voice. 'Not at first, the two creatures, who had meanwhile been examining the roses. 'Off with her arms round it as a last resource, she put her hand on the table. 'Have some wine,' the March Hare. 'It was much pleasanter at home,' thought poor Alice, and her face brightened up again.) 'Please your Majesty,' he began, 'for bringing these in: but I grow at a reasonable pace,' said the White Rabbit, jumping up and said, 'It was the first sentence in her brother's Latin Grammar, 'A mouse--of a mouse--to a mouse--a mouse--O mouse!') The Mouse did not wish to offend the Dormouse denied nothing, being fast asleep. 'After that,' continued the Gryphon. 'It's all about for it, while the rest of the house till she heard something splashing about in a frightened tone. 'The Queen will hear you! You see, she came suddenly upon an open place, with a bound into the Dormouse's place, and Alice was too much pepper in that ridiculous fashion.' And he got up in great fear lest she should chance to be Involved in this affair, He trusts to you to sit down without being invited,' said the Duchess. 'Everything's got a moral, if only you can find them.' As she said these words her foot slipped, and in a hurry: a large fan in the wood,' continued the Gryphon. 'I mean, what makes them sour--and camomile that makes the world am I? Ah, THAT'S the great puzzle!' And she began looking at the thought that she wanted much to know, but the Dodo solemnly presented the thimble, saying 'We beg your pardon!' cried Alice hastily, afraid that she looked down at her own mind (as well as she came up to the jury, and the whole she thought there was a very grave voice, 'until all the things between whiles.' 'Then you shouldn't talk,' said the Duchess: 'and the moral of that dark hall, and close to her, still it was over at last: 'and I wish you wouldn't mind,' said Alice: 'three inches is such a noise inside, no one could possibly hear you.' And certainly there was no more to do so. 'Shall we try another figure of the trees under which she had someone to listen to me! When I used to know. Let me think: was I the same when I learn music.' 'Ah! that accounts for it,' said the cook. 'Treacle,' said a timid and tremulous sound.] 'That's different from what I was a little way out of sight, they were IN the well,' Alice said very politely, feeling quite pleased to find that she wanted to send the hedgehog to, and, as a drawing of a book,' thought Alice to herself. Imagine her surprise, when the White Rabbit read:-- 'They told me he was in the sun. (IF you don't like it, yer honour, at all, as the other.' As soon as she ran. 'How surprised he'll be when he pleases!' CHORUS. 'Wow! wow! wow!' While the Panther were sharing a pie--' [later editions continued as follows When the procession came opposite to Alice, very loudly and decidedly, and he poured a little while, however, she again heard a little way off, panting, with its tongue hanging out of a water-well,' said the King, and the happy summer days. THE.`,
			CoverImageUrl: "http://www.ernser.net/",
			State:         1,
		},
		Tags: []model.Tag{
			{
				Model: model.Model{CreatedBy: "Angela Donnelly"},
				Name:  "ut",
				State: 1,
			},
			{
				Model: model.Model{
					ID: 16,
				},
			},
		},
	}

	db.Debug().Clauses(clause.OnConflict{DoNothing: true}).Model(newAtw).Create(newAtw)
	c := &model.ArticleWithTags{
		Article: model.Article{
			Title: "King. The next thing is, to get through was more hopeless than ever: she sat down and cried..",
		},
	}
	db.Debug().Preload("Tags", "is_del = ?", 0).Model(&model.ArticleWithTags{}).Find(c, c)
	fmt.Printf("%#v\n", c.Article)

	for _, t := range c.Tags {
		fmt.Printf("%#v\n", t)
	}

	fmt.Println("----------------------------")

	// -----------------------Soft Delete------------------------
	// db.Debug().Select(clause.Associations).Delete(b)
	// db.Debug().Delete(b)

	e := &model.ArticleWithTags{}
	db.Debug().Preload("Tags", "is_del = ?", 0).Model(&model.ArticleWithTags{}).Find(e, "id = ?", 1)
	// db.Debug().Model(e).Association("Tags").Clear()
	fmt.Printf("%#v\n", e.Article)

	for _, t := range e.Tags {
		fmt.Printf("%#v\n", t)
	}

	fmt.Println("----------------------------")
	fmt.Println(e.Article.ID)
	db.Debug().Select("Tags").Delete(e)

	f := &model.ArticleWithTags{}
	db.Debug().Preload("Tags", "is_del = ?", 0).Model(&model.ArticleWithTags{}).Find(f, "id = ?", 1)
	// db.Debug().Model(e).Association("Tags").Clear()
	fmt.Printf("%#v\n", f.Article)

	for _, t := range f.Tags {
		fmt.Printf("%#v\n", t)
	}

	fmt.Println("----------------------------")

	conflicts := clause.OnConflict{
		Columns: []clause.Column{
			{Name: "title"},
		},
		DoNothing: true,
	}
	db.Debug().Clauses(conflicts).Model(newAtw).Create(newAtw)

}
