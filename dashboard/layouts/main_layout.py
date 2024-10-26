from config import UPDATE_INTERVAL
from dash import dcc, html


def create_layout() -> html.Div:
    """
    Creates the main layout for the intelligent refrigeration dashboard.

    Returns:
        html.Div: A Dash HTML Div component containing the layout elements, including:
            - A header with the title "Dashboard".
            - A graph component for displaying temperature and humidity data.
            - Buttons to control the fan ("Turn fan on" and "Turn fan off").
            - A div to display the fan status.
            - An interval component for periodic updates.
    """
    return html.Div(
        children=[
            html.H1("Dashboard"),
            dcc.Graph(id="temperature-humidity-graph"),
            html.Div(
                [
                    html.Button("Turn fan on", id="fan-on", n_clicks=0),
                    html.Button("Turn fan off", id="fan-off", n_clicks=0),
                    html.Div(id="fan-status", style={"margin-top": "10px"}),
                ]
            ),
            dcc.Interval(
                id="interval-component", interval=UPDATE_INTERVAL, n_intervals=0
            ),
        ]
    )
