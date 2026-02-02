package de.deutschebahn.rfz.sigsim.config;

/**
 * Demo configuration class for SigSim.
 */
public class SigSimConfig {

    private String simulatorMode;

    public SigSimConfig() {
        this.simulatorMode = "default";
    }

    public String getSimulatorMode() {
        return simulatorMode;
    }

    public void setSimulatorMode(String simulatorMode) {
        this.simulatorMode = simulatorMode;
    }
}
