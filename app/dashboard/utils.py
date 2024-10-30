import os

import pandas as pd
import requests
from dotenv import load_dotenv

load_dotenv(".env")
API_URL = os.getenv("API_URL")

def get_sensor_data() -> pd.DataFrame | None:
    try:
        response = requests.get(f"{API_URL}/temperatures")
        response.raise_for_status()
        data = response.json()

        return pd.DataFrame(data) if data else None
    except requests.RequestException as e:
        print(f"Error fetching sensor data: {e}")
        return None


def control_fan(state: str) -> pd.DataFrame | None:
    try:
        response = requests.post(f"{API_URL}/fan", json={"state": state})
        response.raise_for_status()
        data = response.json()
        return pd.DataFrame([data]) if isinstance(data, dict) else pd.DataFrame(data)
    except requests.RequestException as e:
        print(f"Error controlling fan: {e}")
        return None
