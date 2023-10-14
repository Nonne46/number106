import React, { SetStateAction, useState, useEffect, useMemo } from "react";
import { Renderable, Toast, toast, ValueFunction } from "react-hot-toast";
import TextareaAutosize from "react-textarea-autosize";
import "nes.css/css/nes.min.css";
import {
  GetSpeakers,
  GetPitches,
  GetDevices,
  PlayAudio,
  GetTTS,
  GetEffects,
} from "../wailsjs/go/main/App";

function Radio(props: any) {
  const [items, setItems] = useState<any[]>([]);

  useEffect(() => {
    props
      .getItems()
      .then((res: SetStateAction<any[]>) => {
        setItems(res);
      })
      .catch((err: Renderable | ValueFunction<Renderable, Toast>) =>
        toast.error(err)
      );
  }, [items.length]);

  return (
    <div className="items">
      {items.map((item) =>
        item.speakers ? (
          <div key={item.speakers[0]} className="item">
            <label>
              <input
                type="radio"
                className="nes-radio"
                value={item.speakers[0]}
                name={props.name}
              />
              <span>{item.name}</span>
            </label>
          </div>
        ) : (
          <div key={item} className="item">
            <label>
              <input
                type="radio"
                className="nes-radio"
                value={item}
                name={props.name}
              />
              <span>{item}</span>
            </label>
          </div>
        )
      )}
    </div>
  );
}

function Select(props: any) {
  const [items, setItems] = useState([""]);

  useEffect(() => {
    props
      .getItems()
      .then((res: SetStateAction<string[]>) => setItems(res))
      .catch((err: Renderable | ValueFunction<Renderable, Toast>) =>
        toast.error(err)
      );
  }, [items.length]);

  return (
    <div className="nes-select">
      <select id="default_select">
        <option value="" selected>
          none
        </option>
        {items.map((item) => (
          <option value={item}>{item}</option>
        ))}
      </select>
    </div>
  );
}

async function getSpeakers() {
  let meow = await GetSpeakers();

  meow.voices[0].speakers[0];

  return meow.voices;
}

function App() {
  const [speaker, setSpeaker] = useState("");
  const [effect, setEffect] = useState("");
  // const [pitch, setPitch] = useState("");
  const [device, setDevice] = useState("");
  const [text, setText] = useState("");
  const [prevText, setPrevText] = useState([""]);
  const [token, setToken] = useState("");

  const updateSpeaker = (e: any) => setSpeaker(e.target.value);
  const updateEffect = (e: any) => setEffect(e.target.value);
  // const updatePitch = (e: any) => setPitch(e.target.value);
  const updateDevice = (e: any) => setDevice(e.target.value);
  const updateText = (e: any) => setText(e.target.value);

  const Speakers = useMemo(
    () => <Radio getItems={getSpeakers} name="speaker" />,
    []
  );
  // const Pitches = useMemo(
  //     () => <Radio getItems={GetPitches} name="pitch" />,
  //     []
  // );
  const Devices = useMemo(
    () => <Radio getItems={GetDevices} name="device" />,
    []
  );

  const Effects = useMemo(
    () => <Select getItems={GetEffects} name="effect" />,
    []
  );

  function playAudio() {
    GetTTS(speaker, "medium", text, effect)
      .then((path) => {
        PlayAudio(device, path).catch((err) => toast.error(err));
      })
      .catch((err) => toast.error(err));

    if (text.trim().length != 0) {
      let someText = prevText;
      someText.push(text);
      setPrevText(someText);
    }
    setText("");
  }

  function prev() {
    let someText = prevText;
    setText(someText.pop() as string);
    setPrevText(someText);
  }

  useEffect(() => {
    const handleKeyPress = (e: KeyboardEvent) => {
      if (e.key === "Enter") {
        e.preventDefault();
        playAudio();
      }
    };

    document.addEventListener("keydown", handleKeyPress);

    return () => {
      document.removeEventListener("keydown", handleKeyPress);
    };
  }, [speaker, text, effect, token]);

  return (
    <div id="App">
      <div className="settings">
        <section
          className="nes-container with-title speakers"
          onChange={updateSpeaker}
        >
          <h3 className="title">Speakers</h3>
          {Speakers}
        </section>

        {/* <section
                    className="nes-container with-title pitches"
                    onChange={updatePitch}
                >
                    <h3 className="title">Pitch</h3>
                    {Pitches}
                </section> */}

        <section
          className="nes-container with-title devices"
          onChange={updateDevice}
        >
          <h3 className="title">Output Device</h3>
          {Devices}
        </section>

        <section
          className="nes-container with-title effects"
          onChange={updateEffect}
        >
          <h3 className="title">Effects</h3>
          {Effects}
        </section>
      </div>

      <div className="play">
        <section className="nes-container with-title play">
          <h3 className="title">Озвучь пжлст</h3>
          <TextareaAutosize
            id="textarea_field"
            className="nes-textarea"
            maxLength={1800}
            maxRows={10}
            value={text}
            onChange={updateText}
          />

          <span>
            <button type="submit" className="nes-btn" onClick={playAudio}>
              Play
            </button>

            <button type="button" className="nes-btn" onClick={prev}>
              Prev
            </button>
          </span>
        </section>
      </div>
    </div>
  );
}

export default App;
