Rails.application.routes.draw do
  resources :messages, only: %i[create]
end
