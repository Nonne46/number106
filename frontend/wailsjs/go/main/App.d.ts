// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {tts} from '../models';

export function GetDevices():Promise<Array<string>>;

export function GetEffects():Promise<Array<string>>;

export function GetPitches():Promise<Array<string>>;

export function GetSpeakers():Promise<tts.Speakers>;

export function GetTTS(arg1:string,arg2:string,arg3:string,arg4:string):Promise<string>;

export function IsPlaying():Promise<boolean>;

export function PlayAudio(arg1:string,arg2:string):Promise<void>;
