from dash.dependencies import Input, Output
from services.api_service import control_fan
import dash
def register_fan_callbacks(app: dash.Dash) -> None:
    @app.callback(
        Output("fan-status", "children"),
        [Input("fan-on", "n_clicks"), Input("fan-off", "n_clicks")]
    )
    def update_fan_status(fan_on_clicks, fan_off_clicks):
        ctx = dash.callback_context

        if not ctx.triggered:
            return "Fan Status: Unknown"
        else:
            button_id = ctx.triggered[0]["prop_id"].split(".")[0]
            if button_id == "fan-on":
                message = control_fan("on")
                return f"Status do Ventilador: {message}"
            elif button_id == "fan-off":
                message = control_fan("off")
                return f"Fan Status: {message}"
