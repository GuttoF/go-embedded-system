from dash import Dash
from layouts.main_layout import create_layout
from callbacks.graph_callbacks import register_graph_callbacks
from callbacks.fan_callbacks import register_fan_callbacks

app = Dash(__name__)
app.layout = create_layout()

register_graph_callbacks(app)
register_fan_callbacks(app)

if __name__ == "__main__":
    app.run_server(host="0.0.0.0", port=8050, debug=True)
