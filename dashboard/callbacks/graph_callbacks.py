import plotly.express as px
import plotly.graph_objects as go
from dash import Dash
from dash.dependencies import Input, Output
from services.api_service import fetch_data


def register_graph_callbacks(app: Dash) -> None:
    """
    Registers the callback for updating the temperature and humidity graph.

    Parameters:
    app (dash.Dash): The Dash application instance.

    Callback:
    - Output: Updates the "figure" property of the "temperature-humidity-graph"
    component.
    - Input: Triggers on changes to the "n_intervals" property of the "interval
    component".

    The callback function `update_graph`:
    - Fetches data using the `fetch_data` function.
    - If no data is found, returns an empty line graph with the title
    "There is no data".
    - If data is found, returns a line graph plotting "temperature" and "humidity"
    against "timestamp" with appropriate titles for the graph and axes.

    Returns:
    None
    """

    @app.callback(
        # type: ignore
        Output("temperature-humidity-graph", "figure"),
        [Input("interval-component", "n_intervals")],
    )
    def update_graph(n_intervals: int) -> go.Figure:
        df = fetch_data()
        if df.empty:
            fig = px.line(title="There is no data")
        else:
            fig = px.line(
                df,
                x="timestamp",
                y=["temperature", "humidity"],
                title="Temperature and Humidity",
            )
            fig.update_layout(xaxis_title="Time", yaxis_title="Values")
        return fig

    return None
