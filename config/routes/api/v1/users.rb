namespace :users do
    post '', to: 'users#register'
    get '', to: 'users#get'
end