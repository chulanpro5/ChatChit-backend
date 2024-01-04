from fastapi import FastAPI
from deep_translator import GoogleTranslator

app = FastAPI()

@app.post("/translate/")
async def translate_text(text: str, target_language: str = 'vi'):
    translated = GoogleTranslator(source='auto', target=target_language).translate(text)
    return {"translation": translated}
