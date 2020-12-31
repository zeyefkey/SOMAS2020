type Role = "president" | "judge" | "speaker" | "none"

export class TurnsInRoles {
    president: number;
    judge: number;
    speaker: number;
    none: number;

    constructor(president?: number, judge?: number, speaker?: number, none?: number) {
        this.president = president || 0;
        this.judge = judge || 0;
        this.speaker = speaker || 0;
        this.none = none || 0;
    }

    set(role: Role, val: number) {
        switch(role) {
            case "president": {
                this.president = val;
                break;
            }
            case "judge": {
                this.judge = val;
                break;
            }
            case "speaker": {
                this.speaker = val;
                break;
            }
            case "none": {
                this.none = val;
                break;
            }
        }
    }

    increment(role: Role) {
        switch(role) {
            case "president": {
                this.president += 1;
                break;
            }
            case "judge": {
                this.judge += 1;
                break;
            }
            case "speaker": {
                this.speaker += 1;
                break;
            }
            case "none": {
                this.none += 1;
                break;
            }
        }
    }

    get(role: Role) {
        switch(role) {
            case "president": {
                return this.president;
            }
            case "judge": {
                return this.judge;
            }
            case "speaker": {
                return this.speaker;
            }
            case "none": {
                return this.none;
            }
        }
    }
};

export type ProcessedRoleElement = {
    name: string;
    roles: TurnsInRoles[];
};

export type ProcessedRoleData = ProcessedRoleElement[];
