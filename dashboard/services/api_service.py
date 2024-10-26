import pandas as pd
import requests
from config import API_URL


def fetch_data() -> pd.DataFrame:
    """
    Fetches temperature data from the API and returns it as a pandas DataFrame.

    Returns:
        pd.DataFrame: A DataFrame containing the temperature data with column names in
        lowercase.
                If the API request fails, an empty DataFrame is returned.
    """
    response = requests.get(f"{API_URL}/temperatures")
    if response.status_code == 200:
        data = response.json()
        df = pd.DataFrame(data)
        df.columns = df.columns.str.lower()
        return df
    else:
        return pd.DataFrame()


def control_fan(state: str) -> str:
    """
    Sends a POST request to control the state of a fan.

    Args:
        state (str): The desired state of the fan (e.g., "on", "off").

    Returns:
        str: A message indicating the result of the operation. If the request is
        successful,
            it returns the message from the response JSON or a default success message.
            If the request fails, it returns an error message.
    """
    response = requests.post(f"{API_URL}/fan", json={"state": state})
    if response.status_code == 200:
        message: str = response.json().get("message", "Operation successful")
        return message
    else:
        return "Error: Failed to control fan"
