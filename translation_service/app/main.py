from fastapi import FastAPI
from deep_translator import GoogleTranslator,single_detection

app = FastAPI()

@app.post("/translate/")
async def translate_text(text: str, target_language: str = 'vi'):
    translated = GoogleTranslator(source='auto', target=target_language).translate(text)
    return {"translation": translated}

@app.post("/detect/")
async def detect_language(text: str, api_key: str = None):
    detected = single_detection(text, api_key=api_key)
    return {"language": detected}