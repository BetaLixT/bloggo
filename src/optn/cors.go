package optn

type CorsOptions struct {
  AllowedOrigins []string
}

var _ Options = (*CorsOptions)(nil)

func (coptn *CorsOptions) GetKey() string {
  return "CorsOptions"
}
