import os
import pandas as pd
import plotly.express as px # type: ignore
import plotly.graph_objects as go
import requests
from dash import Dash, dcc, html
from dash.dependencies import Input, Output

API_URL = str(os.getenv("API_URL"))
if not API_URL:
    raise ValueError("The api url is not set")

def fetch_data() -> pd.DataFrame:
    response = requests.get(API_URL)
    if response.status_code == 200:
        data = response.json()
        df = pd.DataFrame(data)
        df.columns = data.columns.str.lower()
        return df
    else:
        return pd.DataFrame()


app = Dash()


app.layout = html.Div(
    children=[
        html.H1("Dashboard"),
        dcc.Interval(id="interval-component", interval=5 * 1000, n_intervals=0),
        dcc.Graph(id="temperature-graph"),
    ]
)


@app.callback(
    # type: ignore
    Output("temperature-graph", "figure"), [Input("interval-component", "n_intervals")]
)
def update_graph(n_intervals: int) -> go.Figure:
    df = fetch_data()

    if df.empty:
        fig = px.line(title="Nenhum dado encontrado")
    else:
        fig = px.line(
            df, x="timestamp", y="temperature", title="Temperaturas ao Longo do Tempo"
        )
        fig.update_layout(xaxis_title="Tempo", yaxis_title="Temperatura (°C)")

    return fig


if __name__ == "__main__":
    app.run_server(host="0.0.0.0", port=8050, debug=True)
