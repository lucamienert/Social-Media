import React, { useState, useEffect } from 'react';
import { useLoading } from '../Contexts/LoadingContext';  // Import your custom loading context
import 'bootstrap/dist/css/bootstrap.min.css';

interface FeedItem {
  id: string;
  title: string;
  content: string;
  date: string;
}

export default function FeedPage() {
  const [feedData, setFeedData] = useState<FeedItem[]>([]);
  const { setLoading } = useLoading();  // Use the setLoading function from context
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    // Set loading to true when data is being fetched
    setLoading(true);

    // Dummy data for now
    const dummyFeed: FeedItem[] = [
      {
        id: '1',
        title: 'First Post',
        content: 'This is the content of the first post. Here, you can share anything you like!',
        date: new Date().toISOString(),
      },
      {
        id: '2',
        title: 'Second Post',
        content: 'Here is the content of the second post. Stay connected with your friends.',
        date: new Date().toISOString(),
      },
      {
        id: '3',
        title: 'Third Post',
        content: 'Content for the third post, it includes some interesting information.',
        date: new Date().toISOString(),
      },
    ];

    setTimeout(() => {
      setFeedData(dummyFeed); // Simulate API delay
      setLoading(false); // Set loading to false when data is fetched
    }, 1500); // Simulating an API call delay of 1.5 seconds

    // Handle error or other actions here
    setLoading(false);
  }); // Ensure useLoading context is included in dependencies

  setLoading(false);

  return (
    <div className="container py-5">
      <h1 className="text-center mb-4">Feed</h1>

      {error && <div className="alert alert-danger">{error}</div>}

      {/* Feed Items */}
      <div className="row">
        {!feedData.length && !error && (
          <div className="col-12 text-center">
            <p>No feed items available.</p>
          </div>
        )}

        {/* Loop over the feedData and display each item in a Bootstrap Card */}
        {feedData.map((item) => (
          <div key={item.id} className="col-12 col-md-6 col-lg-4 mb-4">
            <div className="card shadow-sm">
              <div className="card-body">
                <h5 className="card-title">{item.title}</h5>
                <p className="card-text">{item.content}</p>
                <p className="card-text">
                  <small className="text-muted">{new Date(item.date).toLocaleString()}</small>
                </p>
              </div>
              <div className="card-footer text-muted">
                <button className="btn btn-outline-primary btn-sm">Like</button>
                <button className="btn btn-outline-info btn-sm ms-2">Comment</button>
              </div>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}
