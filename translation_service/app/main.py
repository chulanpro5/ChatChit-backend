from fastapi import FastAPI
from pydantic import BaseModel
from deep_translator import GoogleTranslator,single_detection

app = FastAPI()

class TranslationRequest(BaseModel):
    text: str
    target_language: str = 'vi'

@app.post("/translate/")
async def translate_text(request: TranslationRequest):
    text = request.text
    target_language = request.target_language
    translated = GoogleTranslator(source='auto', target=target_language).translate(text)
    return {"translation": translated}

# @app.post("/detect/")
# async def detect_language(text: str, api_key: str = None):
#     detected = single_detection(text, api_key=api_key)
#     return {"language": detected}