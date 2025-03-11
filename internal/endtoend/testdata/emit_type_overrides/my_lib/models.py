from datetime import date

from pydantic import BaseModel

class Payload(BaseModel):
    name: str
    release_date: date
