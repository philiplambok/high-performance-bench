class MessagesController < ApplicationController
  def create
    render json: Message.create!(message: 'Hello posting from Rails'), status: :ok
  end
end