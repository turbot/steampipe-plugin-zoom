connection "zoom" {
  plugin = "zoom"

  # Zoom API credentials are available to users with Developer role in the account.
  # You need to create a Server-to-Server OAuth app(https://developers.zoom.us/docs/internal-apps/create) to get the credentials.

  # `account_id`(required) - The Zoom account ID.
  # Can also be set with the ZOOM_ACCOUNT_ID environment variable.
  # account_id    = "Xt1aUD4WQ56w7hDhVbtDp"

  # `client_id`(required) - The Zoom Client ID provided by Server-to-Server OAuth app.
  # Can also be set with the ZOOM_CLIENT_ID environment variable.
  # client_id    = "MZw2piRfTsOdpwx2Dh5U"

  # `client_secret`(required) - The Zoom Client Secret provided by Server-to-Server OAuth app.
  # Can also be set with the ZOOM_CLIENT_SECRET environment variable.
  # client_secret    = "04tKwHgFGvwB1M4HPHOBFP0aLHYqUE"
}

