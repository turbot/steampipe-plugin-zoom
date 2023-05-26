connection "zoom" {
  plugin = "zoom"

  # You can authenticate with a Server-to-Server OAuth app or JWT app.
  # If both auth methods are specified, the Server-to-Server OAuth app credentials will be used.
  # We recommend creating a Server-to-Server OAuth app as JWT apps are being deprecated on June 1, 2023 and will be disabled on September 1, 2023 (https://developers.zoom.us/docs/internal-apps/jwt-faq/).

  ## Server-to-Server OAuth app credentials (https://developers.zoom.us/docs/internal-apps/create/)

  # `account_id` (required) - The Zoom account ID.
  # Can also be set with the ZOOM_ACCOUNT_ID environment variable.
  # account_id = "Xt1aUD4WQ56w7hDhVbtDp"

  # `client_id` (required) - The Zoom Client ID provided by Server-to-Server OAuth app.
  # Can also be set with the ZOOM_CLIENT_ID environment variable.
  # client_id = "MZw2piRfTsOdpwx2Dh5U"

  # `client_secret` (required) - The Zoom Client Secret provided by Server-to-Server OAuth app.
  # Can also be set with the ZOOM_CLIENT_SECRET environment variable.
  # client_secret = "04tKwHgFGvwB1M4HPHOBFP0aLHYqUE"

  ## JWT app credentials (https://developers.zoom.us/docs/platform/build/jwt-app/)

  # `api_key` (required) - The Zoom API key provided by JWT OAuth app.
  # Can also be set with the ZOOM_API_KEY environment variable.
  # api_key = "LFMU3oagTjO8_5sYKQVe"

  # `api_secret` (required) - The Zoom API secret provided by JWT OAuth app.
  # Can also be set with the ZOOM_API_SECRET environment variable.
  # api_secret = "PKS96L69nWSFK2y0A07R2k7xGryVbcWiem"
}

