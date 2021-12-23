import React, { useEffect, useState } from "react";
import { getAllVids } from "../services/backend";
import VideoBox from "../components/VideoBox";

const Feed = () => {
  const [videos, setVideos] = useState(["not", "working", "but "]);
  const getVideos = async () => {
    let res = await getAllVids();
    console.log(res);
    setVideos(res);
  };

  useEffect(() => {
    getVideos();
  }, []);

  return (
    <div>
      <h1>Welcome to television</h1>
      <div className="container">
        {videos && videos.map((vid) => <VideoBox video={vid} key={vid.ID} />)}
        {/* definetly needs pagination */}
      </div>
    </div>
  );
};

export default Feed;
