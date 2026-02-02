package de.deutschebahn.rfz.sigsim;

import de.deutschebahn.rfz.sigsim.config.SigSimConfig;

/**
 * Demo main class for SigSim.
 */
public class SigSimApplication {

    public static void main(String[] args) {
        System.out.println("SigSim Simulator starting...");
        SigSimConfig config = new SigSimConfig();
        System.out.println("Simulator Mode: " + config.getSimulatorMode());
    }
}
