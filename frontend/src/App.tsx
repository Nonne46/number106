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
} from "../wailsjs/go/main/App";

function Radio(props: any) {
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
        <div className="items">
            {items.map((item) => (
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
            ))}
        </div>
    );
}

function App() {
    const [speaker, setSpeaker] = useState("");
    const [pitch, setPitch] = useState("");
    const [device, setDevice] = useState("");
    const [text, setText] = useState("");
    const [prevText, setPrevText] = useState("");

    const updateSpeaker = (e: any) => setSpeaker(e.target.value);
    const updatePitch = (e: any) => setPitch(e.target.value);
    const updateDevice = (e: any) => setDevice(e.target.value);

    const Speakers = useMemo(
        () => <Radio getItems={GetSpeakers} name="speaker" />,
        []
    );
    const Pitches = useMemo(
        () => <Radio getItems={GetPitches} name="pitch" />,
        []
    );
    const Devices = useMemo(
        () => <Radio getItems={GetDevices} name="device" />,
        []
    );

    function playAudio() {
        GetTTS(speaker, pitch, text)
            .then((path) => {
                PlayAudio(device, path).catch((err) => toast.error(err));
            })
            .catch((err) => toast.error(err));

        if (text != "")
        {
            setPrevText(text);
        }
        setText("");
    }

    function prev() {
        setText(prevText);
    }

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

                <section
                    className="nes-container with-title pitches"
                    onChange={updatePitch}
                >
                    <h3 className="title">Pitch</h3>
                    {Pitches}
                </section>

                <section
                    className="nes-container with-title devices"
                    onChange={updateDevice}
                >
                    <h3 className="title">Output Device</h3>
                    {Devices}
                </section>
            </div>

            <div className="play">
                <section className="nes-container with-title play">
                    <h3 className="title">Озвучь пжлст</h3>

                    <TextareaAutosize
                        id="textarea_field"
                        className="nes-textarea"
                        maxLength={600}
                        maxRows={10}
                        value={text}
                        onChange={(e) => setText(e.target.value)}
                    />

                    <span>
                        <button
                            type="button"
                            className="nes-btn"
                            onClick={playAudio}
                        >
                            Play
                        </button>

                        <button
                            type="button"
                            className="nes-btn"
                            onClick={prev}
                        >
                            Prev
                        </button>
                    </span>
                </section>
            </div>
        </div>
    );
}

export default App;
