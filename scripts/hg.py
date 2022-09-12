from transformers import AutoTokenizer, AutoModelForSequenceClassification
import numpy as np
import torch
from memory_profiler import profile

@profile
def app():
    tokenizer = AutoTokenizer.from_pretrained("ProsusAI/finbert")

    model = AutoModelForSequenceClassification.from_pretrained("ProsusAI/finbert")
    
    messages=["stocks rallied before closing", "boom and bust cycle of pakisan",
    "End of pakistani terror?", "Is pakistan finally ending its debt problem \
    and working towards success?"]

    tokenized = tokenizer(messages, padding=True, truncation=True, return_tensors="pt")
    outputs = model(**tokenized)
    outputs = torch.nn.functional.softmax(outputs.logits, dim = -1)
    print(outputs)

if __name__ == "__main__":
    app()
