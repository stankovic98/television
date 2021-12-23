export const getAllVids = async () => {
  let videos = await window.backend.GetAllVideos();
  return videos;
};
