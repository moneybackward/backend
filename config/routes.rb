Rails.application.routes.draw do
  # devise_for :users
  mount Rswag::Ui::Engine => '/api-docs'
  mount Rswag::Api::Engine => '/api-docs'

  scope '/api' do
    scope '/v1' do
      resources :users, only: [:create, :index, :show]
    end
  end

  # draw(:routes)
  # Define your application routes per the DSL in https://guides.rubyonrails.org/routing.html

  # Defines the root path route ("/")
  # root "articles#index"
end
