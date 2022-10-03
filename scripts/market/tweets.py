from tweepy.auth import OAuthHandler
import tweepy
from sentiment import sentiment_vader
import json
import nltk
words = set(nltk.corpus.words.words())

access_secret='K5JcjFlD2TOh2csFdZvRKQjF8jBZpW5pioW1JUnQG0FBN'
consumer_key='oen6gdBlEGyySyPDhc4mLY7DI'
consumer_secret= 'EzedL8cENk0nuiG4T6EvwbPH4Pof5lNJTCA0ckVCci2ZPtjsQV'
auth = OAuthHandler(consumer_key, consumer_secret)
auth.set_access_token('203929297-Tl4ed85WNK88ASc4eAgqmXUBbZtwDEB0DPN3L3YD', access_secret)

api = tweepy.API(auth)

public_tweets = api.home_timeline()
# for tweet in public_tweets:
    # print(tweet.text)

sentiment_map = {}
tweet_set = set()
i = 0 
for status in tweepy.Cursor(api.search, "Pakistan",
                            count=1000).items(10000):
    js = status._json
    sent = js['text']
    sent = " ".join(w for w in nltk.wordpunct_tokenize(sent) \
         if w.lower() in words or not w.isalpha())
    tweet_set.add(sent)

print(len(tweet_set))

for sent in tweet_set:
    sentiment_val = sentiment_vader(sent)[4]
    # print(sent, sentiment_val)
    if sentiment_val in sentiment_map:
        sentiment_map[sentiment_val] += 1
    else:
        sentiment_map[sentiment_val] = 1

print(sentiment_map)



# query_map = ["Pakistan Economy", "Nawaz Sharif", "Imran Khan PTI", "Pakistan Dollar Rate"]

# for query in query_map:
#     sentiment_map = {}
#     for tweets in api.search(q=query, lang="en", count=100, result_type='popular'):
#         sentiment_val = sentiment_vader(tweets.text)[4]
#         if sentiment_val in sentiment_map:
#             sentiment_map[sentiment_val] += 1
#         else:
#             sentiment_map[sentiment_val] = 1

#     print(sentiment_map)