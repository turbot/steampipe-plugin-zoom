connection "zoom" {
  plugin = "zoom"

  # Zoom API credentials are available to users with Developer role in the account.
  # You need to create a Server-to-Server OAuth app (https://developers.zoom.us/docs/internal-apps/create) or a SDK/JWT APP (https://marketplace.zoom.us/docs/guides/build/sdk-app) to get the credentials.
  # It is recommended that you create Server-to-Server OAuth since the JWT app has been deprecated as of June 1, 2023 and will be disabled on September 1, 2023. https://developers.zoom.us/docs/internal-apps/jwt-faq/

  # Server-to-Server OAuth app credentials

  # Zoom account ID is required for requests. Required. 
  # This can also be set via the ZOOM_ACCOUNT_ID environment variable.
  # account_id = "Xt1aUD4WQ56w7hDhVbtDp"

  # Zoom Client ID provided by Server-to-Server OAuth app is required for requests. Required.
  # This can also be set via the ZOOM_CLIENT_ID environment variable.
  # client_id = "MZw2piRfTsOdpwx2Dh5U"

  # Zoom Client Secret provided by Server-to-Server OAuth app is required for requests. Required.
  # This can also be set via the ZOOM_CLIENT_SECRET environment variable.
  # client_secret = "04tKwHgFGvwB1M4HPHOBFP0aLHYqUE"

  # SDK/JWT app credentials

  # Zoom API key provided by SDK/JWT OAuth app is required for requests. Required.
  # This can also be set via the ZOOM_API_KEY environment variable.
  # api_key = "LFMU3oagTjO8_5sYKQVe"

  # Zoom API secret provided by SDK/JWT OAuth app is required for requests. Required.
  # This can also be set via the ZOOM_API_SECRET environment variable.
  # api_secret = "PKS96L69nWSFK2y0A07R2k7xGryVbcWiem"

  # If you define Server-to-Server and JWT, then the plugin prioritizes Server-to-Server creds.
}

