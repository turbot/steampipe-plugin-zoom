connection "zoom" {
  plugin = "zoom"

  # Zoom API credentials are available to users with Developer role in the account.
  # You need to create a Server-to-Server OAuth app(https://developers.zoom.us/docs/internal-apps/create) or a SDK/JWT APP (https://marketplace.zoom.us/docs/guides/build/sdk-app)to get the credentials.
  # It is recommended that you create Server-to-Server OAuth as JWT app is deprecated On June 1, 2023 and will be disabled on September 1, 2023. https://developers.zoom.us/docs/internal-apps/jwt-faq/

  ## Server-to-Server OAuth app credentials

  # `account_id`(required) - The Zoom account ID.
  # Can also be set with the ZOOM_ACCOUNT_ID environment variable.
  # account_id = "Xt1aUD4WQ56w7hDhVbtDp"

  # `client_id`(required) - The Zoom Client ID provided by Server-to-Server OAuth app.
  # Can also be set with the ZOOM_CLIENT_ID environment variable.
  # client_id = "MZw2piRfTsOdpwx2Dh5U"

  # `client_secret`(required) - The Zoom Client Secret provided by Server-to-Server OAuth app.
  # Can also be set with the ZOOM_CLIENT_SECRET environment variable.
  # client_secret = "04tKwHgFGvwB1M4HPHOBFP0aLHYqUE"

  ## SDK/JWT app credentials

  # `api_key`(required) - The Zoom API key provided by JWT OAuth app.
  # Can also be set with the ZOOM_API_KEY environment variable.
  # api_key = "LFMU3oagTjO8_5sYKQVe"

  # `api_secret`(required) - The Zoom API secret provided by JWT OAuth app.
  # Can also be set with the ZOOM_API_SECRET environment variable.
  # api_secret = "PKS96L69nWSFK2y0A07R2k7xGryVbcWiem"

  # If you define Server-to-Server and JWT, then the plugin prioritizes Server-to-Server creds.
}

