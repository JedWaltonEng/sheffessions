import React, { ChangeEvent, FormEvent } from 'react';
import Cookies from 'js-cookie';

interface State {
  value: string;
}

class SubmissionForm extends React.Component<{}, State> {
  constructor(props: {}) {
    super(props);
    this.state = { value: '' };

    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleChange(event: ChangeEvent<HTMLTextAreaElement>) {
    this.setState({ value: event.target.value });
  }

  // handleSubmit(event: FormEvent) {
  //   // alert('Your Sheffession has been submitted :) ' + this.state.value);
  //   alert('Your Sheffession has been submitted...');
  //   event.preventDefault();
  // }

  webhookURL = 'https://discord.com/api/webhooks/1141479151178629291/1lFcn3RBk_HizyPNkTXBc95kScxfNoroK4-vzi7xzJZhG5IUuxOOucW2fut7qCQ9qNWO';

  async handleSubmit(event: FormEvent) {
    event.preventDefault();

    const confessionCount = Cookies.get('confessionCount');
    if (confessionCount && parseInt(confessionCount) >= 10) {
      alert('You may only submit 10 confessions daily.');
      return;
    }

    try {
      const response = await fetch(this.webhookURL, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ content: this.state.value })
      });

      if (response.ok) {
        alert('Your Sheffession has been submitted :)');

        // Update the cookie with the new count
        const newCount = confessionCount ? parseInt(confessionCount) + 1 : 1;
        Cookies.set('confessionCount', newCount.toString(), { expires: 1 });
      }
    } catch (error) {
      console.error("There was an error submitting the Sheffession:", error);
      alert('There was an error submitting your Sheffession. Please try again later.');
    }
  }

  render() {
    return (
      <form onSubmit={this.handleSubmit} className="w-full lg:w-3/4 mx-auto">
        <label className="block mb-2 w-full">
          <textarea
            value={this.state.value}
            onChange={this.handleChange}
            className="textarea w-full h-64 p-2 border-2 rounded-md text-lg mb-4 md:h-48 md:p-3 md:text-xl lg:h-64 lg:w-full lg:text-2xl lg:p-4"
            placeholder="Type your anonymous confession here..."
          ></textarea>
        </label>
        <div className="flex justify-center">
          <input type="submit" value="Submit" className="p-2 bg-blue-500 text-white rounded-md" />
        </div>
      </form>
    );
  }


}

export default SubmissionForm;